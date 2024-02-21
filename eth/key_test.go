// description:
// @author renshiwei
// Date: 2024/2/21

package eth

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateEth1Key(t *testing.T) {
	priKey, address, err := CreateEth1Key()
	require.NoError(t, err)

	fmt.Println("priKey:", priKey)
	fmt.Println("address:", address)
}
