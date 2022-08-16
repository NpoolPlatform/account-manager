// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/deposit"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/goodbenefit"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/limitation"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/payment"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/platform"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/schema"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/user"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountMixin := schema.Account{}.Mixin()
	account.Policy = privacy.NewPolicies(accountMixin[0], schema.Account{})
	account.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := account.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	accountMixinFields0 := accountMixin[0].Fields()
	_ = accountMixinFields0
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountMixinFields0[0].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() uint32)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountMixinFields0[1].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() uint32)
	// account.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	account.UpdateDefaultUpdatedAt = accountDescUpdatedAt.UpdateDefault.(func() uint32)
	// accountDescDeletedAt is the schema descriptor for deleted_at field.
	accountDescDeletedAt := accountMixinFields0[2].Descriptor()
	// account.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	account.DefaultDeletedAt = accountDescDeletedAt.Default.(func() uint32)
	// accountDescCoinTypeID is the schema descriptor for coin_type_id field.
	accountDescCoinTypeID := accountFields[1].Descriptor()
	// account.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	account.DefaultCoinTypeID = accountDescCoinTypeID.Default.(func() uuid.UUID)
	// accountDescAddress is the schema descriptor for address field.
	accountDescAddress := accountFields[2].Descriptor()
	// account.DefaultAddress holds the default value on creation for the address field.
	account.DefaultAddress = accountDescAddress.Default.(string)
	// accountDescUsedFor is the schema descriptor for used_for field.
	accountDescUsedFor := accountFields[3].Descriptor()
	// account.DefaultUsedFor holds the default value on creation for the used_for field.
	account.DefaultUsedFor = accountDescUsedFor.Default.(string)
	// accountDescPlatformHoldPrivateKey is the schema descriptor for platform_hold_private_key field.
	accountDescPlatformHoldPrivateKey := accountFields[4].Descriptor()
	// account.DefaultPlatformHoldPrivateKey holds the default value on creation for the platform_hold_private_key field.
	account.DefaultPlatformHoldPrivateKey = accountDescPlatformHoldPrivateKey.Default.(bool)
	// accountDescActive is the schema descriptor for active field.
	accountDescActive := accountFields[5].Descriptor()
	// account.DefaultActive holds the default value on creation for the active field.
	account.DefaultActive = accountDescActive.Default.(bool)
	// accountDescLocked is the schema descriptor for locked field.
	accountDescLocked := accountFields[6].Descriptor()
	// account.DefaultLocked holds the default value on creation for the locked field.
	account.DefaultLocked = accountDescLocked.Default.(bool)
	// accountDescLockedBy is the schema descriptor for locked_by field.
	accountDescLockedBy := accountFields[7].Descriptor()
	// account.DefaultLockedBy holds the default value on creation for the locked_by field.
	account.DefaultLockedBy = accountDescLockedBy.Default.(string)
	// accountDescBlocked is the schema descriptor for blocked field.
	accountDescBlocked := accountFields[8].Descriptor()
	// account.DefaultBlocked holds the default value on creation for the blocked field.
	account.DefaultBlocked = accountDescBlocked.Default.(bool)
	// accountDescID is the schema descriptor for id field.
	accountDescID := accountFields[0].Descriptor()
	// account.DefaultID holds the default value on creation for the id field.
	account.DefaultID = accountDescID.Default.(func() uuid.UUID)
	depositMixin := schema.Deposit{}.Mixin()
	deposit.Policy = privacy.NewPolicies(depositMixin[0], schema.Deposit{})
	deposit.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := deposit.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	depositMixinFields0 := depositMixin[0].Fields()
	_ = depositMixinFields0
	depositFields := schema.Deposit{}.Fields()
	_ = depositFields
	// depositDescCreatedAt is the schema descriptor for created_at field.
	depositDescCreatedAt := depositMixinFields0[0].Descriptor()
	// deposit.DefaultCreatedAt holds the default value on creation for the created_at field.
	deposit.DefaultCreatedAt = depositDescCreatedAt.Default.(func() uint32)
	// depositDescUpdatedAt is the schema descriptor for updated_at field.
	depositDescUpdatedAt := depositMixinFields0[1].Descriptor()
	// deposit.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	deposit.DefaultUpdatedAt = depositDescUpdatedAt.Default.(func() uint32)
	// deposit.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	deposit.UpdateDefaultUpdatedAt = depositDescUpdatedAt.UpdateDefault.(func() uint32)
	// depositDescDeletedAt is the schema descriptor for deleted_at field.
	depositDescDeletedAt := depositMixinFields0[2].Descriptor()
	// deposit.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	deposit.DefaultDeletedAt = depositDescDeletedAt.Default.(func() uint32)
	// depositDescAppID is the schema descriptor for app_id field.
	depositDescAppID := depositFields[1].Descriptor()
	// deposit.DefaultAppID holds the default value on creation for the app_id field.
	deposit.DefaultAppID = depositDescAppID.Default.(func() uuid.UUID)
	// depositDescUserID is the schema descriptor for user_id field.
	depositDescUserID := depositFields[2].Descriptor()
	// deposit.DefaultUserID holds the default value on creation for the user_id field.
	deposit.DefaultUserID = depositDescUserID.Default.(func() uuid.UUID)
	// depositDescCoinTypeID is the schema descriptor for coin_type_id field.
	depositDescCoinTypeID := depositFields[3].Descriptor()
	// deposit.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	deposit.DefaultCoinTypeID = depositDescCoinTypeID.Default.(func() uuid.UUID)
	// depositDescAccountID is the schema descriptor for account_id field.
	depositDescAccountID := depositFields[4].Descriptor()
	// deposit.DefaultAccountID holds the default value on creation for the account_id field.
	deposit.DefaultAccountID = depositDescAccountID.Default.(func() uuid.UUID)
	// depositDescBalance is the schema descriptor for balance field.
	depositDescBalance := depositFields[5].Descriptor()
	// deposit.DefaultBalance holds the default value on creation for the balance field.
	deposit.DefaultBalance = depositDescBalance.Default.(decimal.Decimal)
	// depositDescCollectingTid is the schema descriptor for collecting_tid field.
	depositDescCollectingTid := depositFields[6].Descriptor()
	// deposit.DefaultCollectingTid holds the default value on creation for the collecting_tid field.
	deposit.DefaultCollectingTid = depositDescCollectingTid.Default.(func() uuid.UUID)
	// depositDescID is the schema descriptor for id field.
	depositDescID := depositFields[0].Descriptor()
	// deposit.DefaultID holds the default value on creation for the id field.
	deposit.DefaultID = depositDescID.Default.(func() uuid.UUID)
	goodbenefitMixin := schema.GoodBenefit{}.Mixin()
	goodbenefit.Policy = privacy.NewPolicies(goodbenefitMixin[0], schema.GoodBenefit{})
	goodbenefit.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := goodbenefit.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	goodbenefitMixinFields0 := goodbenefitMixin[0].Fields()
	_ = goodbenefitMixinFields0
	goodbenefitFields := schema.GoodBenefit{}.Fields()
	_ = goodbenefitFields
	// goodbenefitDescCreatedAt is the schema descriptor for created_at field.
	goodbenefitDescCreatedAt := goodbenefitMixinFields0[0].Descriptor()
	// goodbenefit.DefaultCreatedAt holds the default value on creation for the created_at field.
	goodbenefit.DefaultCreatedAt = goodbenefitDescCreatedAt.Default.(func() uint32)
	// goodbenefitDescUpdatedAt is the schema descriptor for updated_at field.
	goodbenefitDescUpdatedAt := goodbenefitMixinFields0[1].Descriptor()
	// goodbenefit.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	goodbenefit.DefaultUpdatedAt = goodbenefitDescUpdatedAt.Default.(func() uint32)
	// goodbenefit.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	goodbenefit.UpdateDefaultUpdatedAt = goodbenefitDescUpdatedAt.UpdateDefault.(func() uint32)
	// goodbenefitDescDeletedAt is the schema descriptor for deleted_at field.
	goodbenefitDescDeletedAt := goodbenefitMixinFields0[2].Descriptor()
	// goodbenefit.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	goodbenefit.DefaultDeletedAt = goodbenefitDescDeletedAt.Default.(func() uint32)
	// goodbenefitDescGoodID is the schema descriptor for good_id field.
	goodbenefitDescGoodID := goodbenefitFields[1].Descriptor()
	// goodbenefit.DefaultGoodID holds the default value on creation for the good_id field.
	goodbenefit.DefaultGoodID = goodbenefitDescGoodID.Default.(func() uuid.UUID)
	// goodbenefitDescAccountID is the schema descriptor for account_id field.
	goodbenefitDescAccountID := goodbenefitFields[2].Descriptor()
	// goodbenefit.DefaultAccountID holds the default value on creation for the account_id field.
	goodbenefit.DefaultAccountID = goodbenefitDescAccountID.Default.(func() uuid.UUID)
	// goodbenefitDescBackup is the schema descriptor for backup field.
	goodbenefitDescBackup := goodbenefitFields[3].Descriptor()
	// goodbenefit.DefaultBackup holds the default value on creation for the backup field.
	goodbenefit.DefaultBackup = goodbenefitDescBackup.Default.(bool)
	// goodbenefitDescID is the schema descriptor for id field.
	goodbenefitDescID := goodbenefitFields[0].Descriptor()
	// goodbenefit.DefaultID holds the default value on creation for the id field.
	goodbenefit.DefaultID = goodbenefitDescID.Default.(func() uuid.UUID)
	limitationMixin := schema.Limitation{}.Mixin()
	limitation.Policy = privacy.NewPolicies(limitationMixin[0], schema.Limitation{})
	limitation.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := limitation.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	limitationMixinFields0 := limitationMixin[0].Fields()
	_ = limitationMixinFields0
	limitationFields := schema.Limitation{}.Fields()
	_ = limitationFields
	// limitationDescCreatedAt is the schema descriptor for created_at field.
	limitationDescCreatedAt := limitationMixinFields0[0].Descriptor()
	// limitation.DefaultCreatedAt holds the default value on creation for the created_at field.
	limitation.DefaultCreatedAt = limitationDescCreatedAt.Default.(func() uint32)
	// limitationDescUpdatedAt is the schema descriptor for updated_at field.
	limitationDescUpdatedAt := limitationMixinFields0[1].Descriptor()
	// limitation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	limitation.DefaultUpdatedAt = limitationDescUpdatedAt.Default.(func() uint32)
	// limitation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	limitation.UpdateDefaultUpdatedAt = limitationDescUpdatedAt.UpdateDefault.(func() uint32)
	// limitationDescDeletedAt is the schema descriptor for deleted_at field.
	limitationDescDeletedAt := limitationMixinFields0[2].Descriptor()
	// limitation.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	limitation.DefaultDeletedAt = limitationDescDeletedAt.Default.(func() uint32)
	// limitationDescCoinTypeID is the schema descriptor for coin_type_id field.
	limitationDescCoinTypeID := limitationFields[1].Descriptor()
	// limitation.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	limitation.DefaultCoinTypeID = limitationDescCoinTypeID.Default.(func() uuid.UUID)
	// limitationDescLimitation is the schema descriptor for limitation field.
	limitationDescLimitation := limitationFields[2].Descriptor()
	// limitation.DefaultLimitation holds the default value on creation for the limitation field.
	limitation.DefaultLimitation = limitationDescLimitation.Default.(string)
	// limitationDescAmount is the schema descriptor for amount field.
	limitationDescAmount := limitationFields[3].Descriptor()
	// limitation.DefaultAmount holds the default value on creation for the amount field.
	limitation.DefaultAmount = limitationDescAmount.Default.(decimal.Decimal)
	// limitationDescID is the schema descriptor for id field.
	limitationDescID := limitationFields[0].Descriptor()
	// limitation.DefaultID holds the default value on creation for the id field.
	limitation.DefaultID = limitationDescID.Default.(func() uuid.UUID)
	paymentMixin := schema.Payment{}.Mixin()
	payment.Policy = privacy.NewPolicies(paymentMixin[0], schema.Payment{})
	payment.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := payment.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	paymentMixinFields0 := paymentMixin[0].Fields()
	_ = paymentMixinFields0
	paymentFields := schema.Payment{}.Fields()
	_ = paymentFields
	// paymentDescCreatedAt is the schema descriptor for created_at field.
	paymentDescCreatedAt := paymentMixinFields0[0].Descriptor()
	// payment.DefaultCreatedAt holds the default value on creation for the created_at field.
	payment.DefaultCreatedAt = paymentDescCreatedAt.Default.(func() uint32)
	// paymentDescUpdatedAt is the schema descriptor for updated_at field.
	paymentDescUpdatedAt := paymentMixinFields0[1].Descriptor()
	// payment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	payment.DefaultUpdatedAt = paymentDescUpdatedAt.Default.(func() uint32)
	// payment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	payment.UpdateDefaultUpdatedAt = paymentDescUpdatedAt.UpdateDefault.(func() uint32)
	// paymentDescDeletedAt is the schema descriptor for deleted_at field.
	paymentDescDeletedAt := paymentMixinFields0[2].Descriptor()
	// payment.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	payment.DefaultDeletedAt = paymentDescDeletedAt.Default.(func() uint32)
	// paymentDescCoinTypeID is the schema descriptor for coin_type_id field.
	paymentDescCoinTypeID := paymentFields[1].Descriptor()
	// payment.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	payment.DefaultCoinTypeID = paymentDescCoinTypeID.Default.(func() uuid.UUID)
	// paymentDescAccountID is the schema descriptor for account_id field.
	paymentDescAccountID := paymentFields[2].Descriptor()
	// payment.DefaultAccountID holds the default value on creation for the account_id field.
	payment.DefaultAccountID = paymentDescAccountID.Default.(func() uuid.UUID)
	// paymentDescCollectingTid is the schema descriptor for collecting_tid field.
	paymentDescCollectingTid := paymentFields[3].Descriptor()
	// payment.DefaultCollectingTid holds the default value on creation for the collecting_tid field.
	payment.DefaultCollectingTid = paymentDescCollectingTid.Default.(func() uuid.UUID)
	// paymentDescAvailableAt is the schema descriptor for available_at field.
	paymentDescAvailableAt := paymentFields[4].Descriptor()
	// payment.DefaultAvailableAt holds the default value on creation for the available_at field.
	payment.DefaultAvailableAt = paymentDescAvailableAt.Default.(func() uint32)
	// paymentDescID is the schema descriptor for id field.
	paymentDescID := paymentFields[0].Descriptor()
	// payment.DefaultID holds the default value on creation for the id field.
	payment.DefaultID = paymentDescID.Default.(func() uuid.UUID)
	platformMixin := schema.Platform{}.Mixin()
	platform.Policy = privacy.NewPolicies(platformMixin[0], schema.Platform{})
	platform.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := platform.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	platformMixinFields0 := platformMixin[0].Fields()
	_ = platformMixinFields0
	platformFields := schema.Platform{}.Fields()
	_ = platformFields
	// platformDescCreatedAt is the schema descriptor for created_at field.
	platformDescCreatedAt := platformMixinFields0[0].Descriptor()
	// platform.DefaultCreatedAt holds the default value on creation for the created_at field.
	platform.DefaultCreatedAt = platformDescCreatedAt.Default.(func() uint32)
	// platformDescUpdatedAt is the schema descriptor for updated_at field.
	platformDescUpdatedAt := platformMixinFields0[1].Descriptor()
	// platform.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	platform.DefaultUpdatedAt = platformDescUpdatedAt.Default.(func() uint32)
	// platform.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	platform.UpdateDefaultUpdatedAt = platformDescUpdatedAt.UpdateDefault.(func() uint32)
	// platformDescDeletedAt is the schema descriptor for deleted_at field.
	platformDescDeletedAt := platformMixinFields0[2].Descriptor()
	// platform.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	platform.DefaultDeletedAt = platformDescDeletedAt.Default.(func() uint32)
	// platformDescCoinTypeID is the schema descriptor for coin_type_id field.
	platformDescCoinTypeID := platformFields[1].Descriptor()
	// platform.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	platform.DefaultCoinTypeID = platformDescCoinTypeID.Default.(func() uuid.UUID)
	// platformDescAccountID is the schema descriptor for account_id field.
	platformDescAccountID := platformFields[2].Descriptor()
	// platform.DefaultAccountID holds the default value on creation for the account_id field.
	platform.DefaultAccountID = platformDescAccountID.Default.(func() uuid.UUID)
	// platformDescUsedFor is the schema descriptor for used_for field.
	platformDescUsedFor := platformFields[3].Descriptor()
	// platform.DefaultUsedFor holds the default value on creation for the used_for field.
	platform.DefaultUsedFor = platformDescUsedFor.Default.(string)
	// platformDescBackup is the schema descriptor for backup field.
	platformDescBackup := platformFields[4].Descriptor()
	// platform.DefaultBackup holds the default value on creation for the backup field.
	platform.DefaultBackup = platformDescBackup.Default.(bool)
	// platformDescID is the schema descriptor for id field.
	platformDescID := platformFields[0].Descriptor()
	// platform.DefaultID holds the default value on creation for the id field.
	platform.DefaultID = platformDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	user.Policy = privacy.NewPolicies(userMixin[0], schema.User{})
	user.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := user.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() uint32)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() uint32)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() uint32)
	// userDescDeletedAt is the schema descriptor for deleted_at field.
	userDescDeletedAt := userMixinFields0[2].Descriptor()
	// user.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	user.DefaultDeletedAt = userDescDeletedAt.Default.(func() uint32)
	// userDescAppID is the schema descriptor for app_id field.
	userDescAppID := userFields[1].Descriptor()
	// user.DefaultAppID holds the default value on creation for the app_id field.
	user.DefaultAppID = userDescAppID.Default.(func() uuid.UUID)
	// userDescUserID is the schema descriptor for user_id field.
	userDescUserID := userFields[2].Descriptor()
	// user.DefaultUserID holds the default value on creation for the user_id field.
	user.DefaultUserID = userDescUserID.Default.(func() uuid.UUID)
	// userDescCoinTypeID is the schema descriptor for coin_type_id field.
	userDescCoinTypeID := userFields[3].Descriptor()
	// user.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	user.DefaultCoinTypeID = userDescCoinTypeID.Default.(func() uuid.UUID)
	// userDescAccountID is the schema descriptor for account_id field.
	userDescAccountID := userFields[4].Descriptor()
	// user.DefaultAccountID holds the default value on creation for the account_id field.
	user.DefaultAccountID = userDescAccountID.Default.(func() uuid.UUID)
	// userDescUsedFor is the schema descriptor for used_for field.
	userDescUsedFor := userFields[5].Descriptor()
	// user.DefaultUsedFor holds the default value on creation for the used_for field.
	user.DefaultUsedFor = userDescUsedFor.Default.(string)
	// userDescLabels is the schema descriptor for labels field.
	userDescLabels := userFields[6].Descriptor()
	// user.DefaultLabels holds the default value on creation for the labels field.
	user.DefaultLabels = userDescLabels.Default.([]string)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.10.1"                                         // Version of ent codegen.
	Sum     = "h1:dM5h4Zk6yHGIgw4dCqVzGw3nWgpGYJiV4/kyHEF6PFo=" // Sum of ent codegen.
)
