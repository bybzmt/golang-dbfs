# golang-dbfs

这是一个远程文件读写操作库。

程序实现了os库中所有对文件操作的API，并且接口定义完全一至。

主要是用数据库来进行文件名称的存储，这样子让文件分布到不同的服务器上

用到mysql和memcache

	CREATE TABLE IF NOT EXISTS `node` (
	  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	  `pid` int(11) unsigned NOT NULL,
	  `name` varchar(255) COLLATE latin1_bin NOT NULL,
	  `type` tinyint(4) unsigned NOT NULL,
	  `ctime` int(11) unsigned NOT NULL,
	  `mtime` int(11) unsigned NOT NULL
	  PRIMARY KEY (`id`),
	  UNIQUE KEY `path` (`pid`,`name`);
	) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin COMMENT='文件节点表';

	CREATE TABLE IF NOT EXISTS `file` (
	  `nid` int(10) unsigned NOT NULL AUTO_INCREMENT,
	  `sid` smallint(5) unsigned NOT NULL,
	  `file` varchar(255) COLLATE latin1_bin NOT NULL
	  KEY `nid` (`nid`);
	) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin COMMENT='文件位置';

	CREATE TABLE IF NOT EXISTS `storage` (
	  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	  `host` varchar(255) COLLATE latin1_bin NOT NULL,
	  `port` int(11) NOT NULL,
	  `path` varchar(255) COLLATE latin1_bin NOT NULL COMMENT '匹配路径'
	  PRIMARY KEY (`id`);
	) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin COMMENT='存储服器';


memcache 缓存key： `pid + "/" + name`
