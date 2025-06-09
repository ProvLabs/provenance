# Vault Concepts

The `x/vault` module provides a framework for creating tokenized vaults with advanced features for controlling the deposit and withdrawal of funds.

**Vaults** are based on the [x/marker](/x/marker/spec/README.md) module. They act as a controlling layer over a specific marker, managing deposits of the marker's underlying asset. This is coupled with a mint/burn function that tracks these deposits by issuing the marker's own restricted coin as "shares". The module's design is heavily inspired by the **ERC-4626 Tokenized Vault Standard**, providing a familiar and standardized interface.

---

## Vaults

A **Vault** is an abstraction layer that adds enhanced functionality to a `marker`. It manages a pool of a single underlying asset defined by its associated marker. The core mechanism involves users **depositing** the asset into the vault, for which the vault then **mints** the marker's restricted coin back to the user as **shares**. Conversely, users can **withdraw** the underlying asset by **burning** their shares.

For a vault to function, its module account **must be granted** `MINT`, `BURN`, `DEPOSIT`, and `WITHDRAW` permissions on the associated `marker`. A vault may also inherit and enforce other on-chain restrictions from its marker, such as allow/deny lists.

Each vault is highly configurable at instantiation via the [MsgCreateVaultRequest](03_messages.md#msgcreatevaultrequest) message, allowing for fine-tuned control over its behavior. The primary features include:

* **Fees**
    * **Deposit & Withdraw:** Vaults can charge fees on entry or exit.
    * **Fee Schedules:** Fees can be a simple flat rate or be tiered based on the transaction amount, calculated using either a fixed coin value or a basis point (BPS) percentage.

* **Limits**
    * **Deposit Limits:** Can be configured with an absolute maximum vault size and/or a rolling limit over a given time period (e.g., max 1M per 24 hours).
    * **Withdrawal Limits:** Can be restricted by a total rolling limit for the vault, a per-user rolling limit, and time locks to enforce a holding period.

* **Operational Controls**
    * **Pausing:** Deposits and withdrawals can be independently paused and resumed by the vault's administrator, allowing for maintenance or risk management.

The vault's interface mirrors the core functions of the ERC-4626 standard, offering familiar actions like `deposit`, `withdraw`, and `redeem`. It also provides powerful query endpoints that allow front-ends to preview transactions (e.g., `previewDeposit`, `previewRedeem`) and check limits (`maxWithdraw`) before a user commits to a transaction.