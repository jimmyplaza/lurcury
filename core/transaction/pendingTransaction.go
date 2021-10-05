package transaction

import (
	"lurcury/types"
)

func DeletPendingTransaction(core_arg *types.CoreStruct, i int)(*types.CoreStruct){
	//fmt.Println(core_arg)
	core_arg.PendingTransaction[i] = core_arg.PendingTransaction[len(core_arg.PendingTransaction)-1]
	core_arg.PendingTransaction = core_arg.PendingTransaction[:len(core_arg.PendingTransaction)-1]
	//fmt.Println(core_arg)
	return core_arg
}

func OrderDeletPendingTransaction(core_arg *types.CoreStruct, i int)(*types.CoreStruct){
	core_arg.PendingTransaction = core_arg.PendingTransaction[:i+copy(core_arg.PendingTransaction[i:], core_arg.PendingTransaction[i+1:])]
    return core_arg
}
