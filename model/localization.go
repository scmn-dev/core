package model

// Language ...
type Language struct {
	Categories           string `yaml:"categories" json:"categories"`
	Logins               string `yaml:"logins" json:"logins"`
	CreditCards          string `yaml:"credit_cards" json:"credit_cards"`
	Notes                string `yaml:"notes" json:"notes"`
	Emails               string `yaml:"emails" json:"emails"`
	Identities           string `yaml:"identities" json:"identities"`
	LicenseKeys          string `yaml:"license_keys" json:"license_keys"`
	AllItems             string `yaml:"all_items" json:"all_items"`
	Favorites            string `yaml:"favorites" json:"favorites"`
	Trash                string `yaml:"trash" json:"trash"`
	Server               string `yaml:"server" json:"server"`
	Host                 string `yaml:"host" json:"host"`
	Port                 string `yaml:"port" json:"port"`
	Protocol             string `yaml:"protocol" json:"protocol"`
	LicenseKey           string `yaml:"license_key" json:"license_key"`
	Key                  string `yaml:"key" json:"key"`
	URL                  string `yaml:"url" json:"url"`
	BankName             string `yaml:"bank_name" json:"bank_name"`
	BankCode             string `yaml:"bank_code" json:"bank_code"`
	AccountName          string `yaml:"account_name" json:"account_name"`
	AccountNumber        string `yaml:"account_number" json:"account_number"`
	Iban                 string `yaml:"iban" json:"iban"`
	Currency             string `yaml:"currency" json:"currency"`
	CardName             string `yaml:"card_name" json:"card_name"`
	CardholderName       string `yaml:"cardholder_name" json:"cardholder_name"`
	Type                 string `yaml:"type" json:"type"`
	Number               string `yaml:"number" json:"number"`
	VerificationNumber   string `yaml:"verification_number" json:"verification_number"`
	ExpiryDate           string `yaml:"expiry_date" json:"expiry_date"`
	Note                 string `yaml:"note" json:"note"`
	Export               string `yaml:"export" json:"export"`
	Import               string `yaml:"import" json:"import"`
	Backup               string `yaml:"backup" json:"backup"`
	Restore              string `yaml:"restore" json:"restore"`
	SelectBackup         string `yaml:"select_backup" json:"select_backup"`
	BackupName           string `yaml:"backup_name" json:"backup_name"`
	Email                string `yaml:"email" json:"email"`
	Username             string `yaml:"username" json:"username"`
	Password             string `yaml:"password" json:"password"`
	MasterPassword       string `yaml:"master_password" json:"master_password"`
	MasterPasswordAlert  string `yaml:"master_password_alert" json:"master_password_alert"`
	Signin               string `yaml:"signin" json:"signin"`
	Logout               string `yaml:"logout" json:"logout"`
	BaseURL              string `yaml:"base_url" json:"base_url"`
	SignToDashboard      string `yaml:"sign_to_dashboard" json:"sign_to_dashboard"`
	UseThis              string `yaml:"use_this" json:"use_this"`
	Save                 string `yaml:"save" json:"save"`
	Cancel               string `yaml:"cancel" json:"cancel"`
	Yes                  string `yaml:"yes" json:"yes"`
	No                   string `yaml:"no" json:"no"`
	Refresh              string `yaml:"refresh" json:"refresh"`
	Copy                 string `yaml:"copy" json:"copy"`
	Show                 string `yaml:"show" json:"show"`
	Hide                 string `yaml:"hide" json:"hide"`
	Delete               string `yaml:"delete" json:"delete"`
	Update               string `yaml:"update" json:"update"`
	New                  string `yaml:"new" json:"new"`
	Required             string `yaml:"required" json:"required"`
	AreYouSure           string `yaml:"are_you_sure" json:"are_you_sure"`
	AutoGeneratePassword string `yaml:"auto_generate_password" json:"auto_generate_password"`
	UsedPassword         string `yaml:"used_password" json:"used_password"`
	ConfirmUsedPassword  string `yaml:"confirm_used_password" json:"confirm_used_password"`
	CreatedAt            string `yaml:"created_at" json:"created_at"`
	UpdatedAt            string `yaml:"updated_at" json:"updated_at"`
	DeletedAt            string `yaml:"deleted_at" json:"deleted_at"`
	NetworkError         string `yaml:"network_error" json:"network_error"`
	Error                string `yaml:"error" json:"error"`
	Success              string `yaml:"success" json:"success"`
}
