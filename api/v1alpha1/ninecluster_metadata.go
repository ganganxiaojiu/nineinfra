package v1alpha1

var NineDatahouseClusterset = []ClusterInfo{
	{
		Type:    KyuubiClusterType,
		Version: "v1.8.1",
		Resource: ResourceConfig{
			Replicas: 1,
		},
		Configs: ClusterConfig{
			Image: ImageConfig{
				Repository: "nineinfra/kyuubi",
				Tag:        "v1.8.1-minio",
				PullPolicy: "IfNotPresent",
			},
			Conf: map[string]string{
				"kyuubi.frontend.connection.url.use.hostname": "false",
				"kyuubi.frontend.thrift.binary.bind.port":     "10009",
				"kyuubi.frontend.thrift.http.bind.port":       "10010",
				"kyuubi.frontend.rest.bind.port":              "10099",
				"kyuubi.frontend.mysql.bind.port":             "3309",
				"kyuubi.frontend.protocols":                   "REST,THRIFT_BINARY",
				"kyuubi.metrics.enabled":                      "false",
			},
		},
		ClusterRefs: []ClusterType{
			SparkClusterType,
		},
	},
	{
		Type:    MetaStoreClusterType,
		Version: "v3.1.3",
		Configs: ClusterConfig{
			Image: ImageConfig{
				Repository: "nineinfra/metastore",
				Tag:        "v3.1.3",
				PullPolicy: "IfNotPresent",
			},
			Conf: map[string]string{
				"hive.metastore.warehouse.dir": DataHouseDir,
			},
		},
	},
	{
		Type:    MinioClusterType,
		Version: "RELEASE.2023-09-07T02-05-02Z",
		Configs: ClusterConfig{
			Image: ImageConfig{
				Repository: "minio/minio",
				Tag:        "RELEASE.2023-09-07T02-05-02Z",
				PullPolicy: "IfNotPresent",
			},
		},
	},
	{
		Type:    DatabaseClusterType,
		SubType: DbTypePostgres,
		Version: "v15.2.0",
	},
	{
		Type:    SparkClusterType,
		Version: "v3.2.4",
		Configs: ClusterConfig{
			Image: ImageConfig{
				Repository: "nineinfra/spark",
				Tag:        "v3.2.4-minio",
				PullPolicy: "IfNotPresent",
			},
		},
	},
}
