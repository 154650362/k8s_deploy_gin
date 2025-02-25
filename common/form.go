package common

import "time"

type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=16"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
}

type RegisterForm struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=16"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=1,max=16"`
}

type ResetPassForm struct {
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
}

type EditForm struct {
	Role uint `form:"role" json:"role" binding:"required,oneof=1 2 3"`
}

type UpdateForm struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=1,max=16"` // 昵称
	Avatar   string `form:"avatar" json:"avatar"`
}
type NsAddForm struct {
	Uid         uint       `form:"u_id" json:"u_id" binding:"gte=0"`
	Name        string     `form:"name" json:"name" binding:"required,min=3,max=16"`
	ExpiredTime *time.Time `form:"expired_time" json:"expired_time"`
	Cpu         string     `json:"cpu" form:"cpu" binding:"required"`
	Memory      string     `json:"memory" form:"memory" binding:"required"`
	Num         int        `json:"num" form:"num" binding:"required"`
}

type SparkAddForm struct {
	Uid            uint       `form:"u_id" json:"u_id" binding:"required,gt=0"`
	MasterReplicas int32      `form:"master_replicas" json:"master_replicas" binding:"required,gte=1,lte=3"`
	WorkerReplicas int32      `form:"worker_replicas" json:"worker_replicas" binding:"required,gte=2,lte=10"`
	ExpiredTime    *time.Time `form:"expired_time" json:"expired_time"`
	Cpu            string     `json:"cpu" form:"cpu" binding:"required"`
	Memory         string     `json:"memory" form:"memory" binding:"required"`
}

// BatchSparkAddForm 支持批量添加
type BatchSparkAddForm struct {
	Uid            []uint     `form:"u_id" json:"u_id" binding:"required"`
	MasterReplicas int32      `form:"master_replicas" json:"master_replicas" binding:"required,gte=1,lte=3"`
	WorkerReplicas int32      `form:"worker_replicas" json:"worker_replicas" binding:"required,gte=2,lte=10"`
	ExpiredTime    *time.Time `form:"expired_time" json:"expired_time"`
	Cpu            string     `json:"cpu" form:"cpu" binding:"required"`
	Memory         string     `json:"memory" form:"memory" binding:"required"`
}
type SparkUpdateForm struct {
	Name           string     `json:"name" form:"name" binding:"required"`
	Uid            uint       `form:"u_id" json:"u_id" binding:"required,gt=0"`
	MasterReplicas int32      `form:"master_replicas" json:"master_replicas" binding:"required,gte=1,lte=3"`
	WorkerReplicas int32      `form:"worker_replicas" json:"worker_replicas" binding:"required,gte=2,lte=10"`
	ExpiredTime    *time.Time `form:"expired_time" json:"expired_time"`
	Cpu            string     `json:"cpu" form:"cpu" binding:"required"`
	Memory         string     `json:"memory" form:"memory" binding:"required"`
}
type LinuxAddForm struct {
	Uid         uint       `form:"u_id" json:"u_id" binding:"required,gt=0"`
	Kind        uint       `form:"kind" json:"kind" binding:"required,gt=0"`
	ExpiredTime *time.Time `form:"expired_time" json:"expired_time"`
	Cpu         string     `json:"cpu" form:"cpu" binding:"required"`
	Memory      string     `json:"memory" form:"memory" binding:"required"`
}
type BatchLinuxAddForm struct {
	Uid         []uint     `form:"u_id" json:"u_id" binding:"required"`
	Kind        uint       `form:"kind" json:"kind" binding:"required,gt=0"`
	ExpiredTime *time.Time `form:"expired_time" json:"expired_time"`
	Cpu         string     `json:"cpu" form:"cpu" binding:"required"`
	Memory      string     `json:"memory" form:"memory" binding:"required"`
}
type LinuxUpdateForm struct {
	Name        string     `json:"name" form:"name" binding:"required"`
	Uid         uint       `form:"u_id" json:"u_id" binding:"required,gt=0"`
	ExpiredTime *time.Time `form:"expired_time" json:"expired_time"`
	Cpu         string     `json:"cpu" form:"cpu" binding:"required"`
	Memory      string     `json:"memory" form:"memory" binding:"required"`
}

type HadoopAddForm struct {
	Uid                uint       `form:"u_id" json:"u_id" binding:"required,gt=0"`
	HdfsMasterReplicas int32      `form:"hdfs_master_replicas" json:"hdfs_master_replicas" binding:"required,gte=1,lte=3"`
	DatanodeReplicas   int32      `form:"datanode_replicas" json:"datanode_replicas" binding:"required,gte=2,lte=10"`
	YarnMasterReplicas int32      `form:"yarn_master_replicas" json:"yarn_master_replicas" binding:"required,gte=1,lte=3"`
	YarnNodeReplicas   int32      `form:"yarn_node_replicas" json:"yarn_node_replicas" binding:"required,gte=2,lte=10"`
	ExpiredTime        *time.Time `form:"expired_time" json:"expired_time"`
	Cpu                string     `json:"cpu" form:"cpu" binding:"required"`
	Memory             string     `json:"memory" form:"memory" binding:"required"`
}

// BatchHadoopAddForm 批量添加
type BatchHadoopAddForm struct {
	Uid                []uint     `form:"u_id" json:"u_id" binding:"required"`
	HdfsMasterReplicas int32      `form:"hdfs_master_replicas" json:"hdfs_master_replicas" binding:"required,gte=1,lte=3"`
	DatanodeReplicas   int32      `form:"datanode_replicas" json:"datanode_replicas" binding:"required,gte=2,lte=10"`
	YarnMasterReplicas int32      `form:"yarn_master_replicas" json:"yarn_master_replicas" binding:"required,gte=1,lte=3"`
	YarnNodeReplicas   int32      `form:"yarn_node_replicas" json:"yarn_node_replicas" binding:"required,gte=2,lte=10"`
	ExpiredTime        *time.Time `form:"expired_time" json:"expired_time"`
	Cpu                string     `json:"cpu" form:"cpu" binding:"required"`
	Memory             string     `json:"memory" form:"memory" binding:"required"`
}

type HadoopUpdateForm struct {
	Name               string     `json:"name" form:"name" binding:"required"`
	Uid                uint       `form:"u_id" json:"u_id" binding:"required,gt=0"`
	HdfsMasterReplicas int32      `form:"hdfs_master_replicas" json:"hdfs_master_replicas" binding:"required,gte=1,lte=3"`
	DatanodeReplicas   int32      `form:"datanode_replicas" json:"datanode_replicas" binding:"required,gte=2,lte=10"`
	YarnMasterReplicas int32      `form:"yarn_master_replicas" json:"yarn_master_replicas" binding:"required,gte=1,lte=3"`
	YarnNodeReplicas   int32      `form:"yarn_node_replicas" json:"yarn_node_replicas" binding:"required,gte=2,lte=10"`
	ExpiredTime        *time.Time `form:"expired_time" json:"expired_time"`
	Cpu                string     `json:"cpu" form:"cpu" binding:"required"`
	Memory             string     `json:"memory" form:"memory" binding:"required"`
}

type YamlApplyForm struct {
	Kind string      `json:"kind" form:"kind"`
	Name string      `json:"name" form:"name"`
	Ns   string      `json:"ns" form:"ns"`
	Yaml interface{} `json:"yaml" form:"yaml" binding:"required"`
}

type YamlCreateForm struct {
	Kind string      `json:"kind" form:"kind"`
	Ns   string      `json:"ns" form:"ns"`
	Yaml interface{} `json:"yaml" form:"yaml" binding:"required"`
}

type DeployAddForm struct {
	Name      string `json:"name" form:"name"`
	Namespace string `json:"namespace" form:"namespace"`
	Uid       uint   `json:"u_id" form:"u_id"`
}

type NodeAddForm struct {
	Nodes []struct {
		Host string `json:"host" form:"host" binding:"required"`
	} `json:"nodes" form:"nodes" binding:"required"`

	Port     int    `json:"port" form:"port" binding:"required"`
	User     string `json:"user" form:"user" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
