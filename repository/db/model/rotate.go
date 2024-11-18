package model

import "github.com/jinzhu/gorm"

/*
CREATE TABLE `rotate` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `pid` int(10) unsigned NOT NULL COMMENT '商品id',
 `img_path` varchar(40) NOT NULL COMMENT '图片地址',
 `expire_time` timestamp NOT NULL COMMENT 'token失效时间',
 `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 PRIMARY KEY (`id`),
 UNIQUE KEY `uniq_uid` (`uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='首页轮播图表';
*/

type Rotate struct {
	gorm.Model
	Pid      uint64
	ImagPath string
}

func (r *Rotate) TableName() string {
	return "rotate"
}
