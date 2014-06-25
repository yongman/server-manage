package commit

//分配结果持久化存储相关操作

import (
	"../../modules/alloc"
	"../../modules/db"
	cdao "../../modules/db/commit"
	mdao "../../modules/db/machine"
	"../../modules/log"
	"../../utils"
	"fmt"
)

func DoCommit(host string, port int32, allocBox string) string {
	uuid := utils.GenerateUUID()
	//log.Info(uuid, host, port, allocBox)
	commit := db.Commit{}
	commit.CommitID = uuid
	commit.Host = host
	commit.Port = port
	commit.AllocBox = allocBox
	//调用DAO保存数据
	cdao.AddCommit(&commit)
	//更新主机信息
	err := alloc.DecBoxOne(host, allocBox)
	if err != nil {
		log.Info(err)
	}
	return uuid
}

func DropCommit(commitID string) (err error) {
	if commitID == "all" {
		utils.DangerConfirm("Drop all commits")
		cids, err := cdao.GetAllCommitID()
		if err != nil {
			return err
		}
		for _, cid := range *cids {
			DropCommit(cid)
		}
		return nil
	} else {
		host, boxtype, err := cdao.GetHostBoxByID(commitID)
		if err != nil {
			return err
		}
		err = cdao.DropCommit(commitID)
		if err != nil {
			return err
		}
		log.Info("Dropping commit", commitID)
		//将该条提交撤销，将box添加到原来的主机
		m := mdao.GetMachineByHost(host) //获取该主机信息
		if m != nil {
			if boxtype == utils.Box10G {
				m.Mem.Box10G++
			} else if boxtype == utils.Box5G {
				m.Mem.Box5G++
			} else if boxtype == utils.Box1G {
				m.Mem.Box1G++
			}
			mdao.UpdateMachineMem(m)
		}
		return nil
	}
}

//打印最新的n条提交
func ListCommit(n int) {
	res := cdao.QueryLatestCommit(n)
	//打印
	for _, r := range *res {
		fmt.Println(r)
	}
}

/*
func DropAllCommits() error {
	utils.DangerConfirm("Drop all commits")
	return cdao.DropAll()
}
*/
