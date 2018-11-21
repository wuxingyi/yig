package crypto

import (
	"github.com/journeymidnight/yig/helper"
)

func NewKMS() KMS {
	switch helper.CONFIG.KMSType {
	case "vault":
		c, err := NewVaultConfig(helper.CONFIG.KMSConfigPath)
		if err != nil {
			panic("read kms vault err:" + err.Error())
		}
		vault, err := NewVault(c)
		if err != nil {
			panic("create vault err:" + err.Error())
		}
		return vault

	//extention case here

	default:
		helper.Logger.Println(5, "not support kms type", helper.CONFIG.KMSType)
		return nil
	}
}
