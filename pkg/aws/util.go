package aws

import (
	"airdrop-bot/log"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func WriteAwsConfig(key, secret, dir string) error {
	ss := `[default]
aws_access_key_id = %s
aws_secret_access_key = %s
`
	os.MkdirAll(path.Dir(dir), 0755)

	log.Infof("write aws config to dir %v", dir)
	ss = fmt.Sprintf(ss, key, secret)
	return ioutil.WriteFile(dir, []byte(ss), 0655)
}
