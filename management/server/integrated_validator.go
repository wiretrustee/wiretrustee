package server

import (
	"context"
	"errors"

	nbgroup "github.com/netbirdio/netbird/management/server/group"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
	log "github.com/sirupsen/logrus"

	"github.com/netbirdio/netbird/management/server/account"
)

// UpdateIntegratedValidatorGroups updates the integrated validator groups for a specified account.
// It retrieves the account associated with the provided userID, then updates the integrated validator groups
// with the provided list of group ids. The updated account is then saved.
//
// Parameters:
//   - accountID: The ID of the account for which integrated validator groups are to be updated.
//   - userID: The ID of the user whose account is being updated.
//   - groups: A slice of strings representing the ids of integrated validator groups to be updated.
//
// Returns:
//   - error: An error if any occurred during the process, otherwise returns nil
func (am *DefaultAccountManager) UpdateIntegratedValidatorGroups(ctx context.Context, accountID string, userID string, groups []string) error {
	ok, err := am.GroupValidation(ctx, accountID, groups)
	if err != nil {
		log.WithContext(ctx).Debugf("error validating groups: %s", err.Error())
		return err
	}

	if !ok {
		log.WithContext(ctx).Debugf("invalid groups")
		return errors.New("invalid groups")
	}

	unlock := am.Store.AcquireWriteLockByUID(ctx, accountID)
	defer unlock()

	a, err := am.Store.GetAccountByUser(ctx, userID)
	if err != nil {
		return err
	}

	var extra *account.ExtraSettings

	if a.Settings.Extra != nil {
		extra = a.Settings.Extra
	} else {
		extra = &account.ExtraSettings{}
		a.Settings.Extra = extra
	}
	extra.IntegratedValidatorGroups = groups
	return am.Store.SaveAccount(ctx, a)
}

func (am *DefaultAccountManager) GroupValidation(ctx context.Context, accountID string, groupIDs []string) (bool, error) {
	if len(groupIDs) == 0 {
		return true, nil
	}

	err := am.Store.ExecuteInTransaction(ctx, func(transaction Store) error {
		for _, groupID := range groupIDs {
			_, err := transaction.GetGroupByID(context.Background(), LockingStrengthShare, accountID, groupID)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (am *DefaultAccountManager) GetValidatedPeers(ctx context.Context, accountID string) (map[string]struct{}, error) {
	var err error
	var groups []*nbgroup.Group
	var peers []*nbpeer.Peer
	var settings *Settings

	err = am.Store.ExecuteInTransaction(ctx, func(transaction Store) error {
		groups, err = transaction.GetAccountGroups(ctx, LockingStrengthShare, accountID)
		if err != nil {
			return err
		}

		peers, err = transaction.GetAccountPeers(ctx, LockingStrengthShare, accountID)
		if err != nil {
			return err
		}

		settings, err = transaction.GetAccountSettings(ctx, LockingStrengthShare, accountID)
		return err
	})
	if err != nil {
		return nil, err
	}

	groupsMap := make(map[string]*nbgroup.Group, len(groups))
	for _, group := range groups {
		groupsMap[group.ID] = group
	}

	peersMap := make(map[string]*nbpeer.Peer, len(peers))
	for _, peer := range peers {
		peersMap[peer.ID] = peer
	}

	return am.integratedPeerValidator.GetValidatedPeers(accountID, groupsMap, peersMap, settings.Extra)
}
