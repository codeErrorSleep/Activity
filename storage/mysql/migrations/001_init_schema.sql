-- 活动表
CREATE TABLE IF NOT EXISTS activities (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    category VARCHAR(50) NOT NULL COMMENT '活动类型',
    version VARCHAR(20) NOT NULL COMMENT '活动版本',
    name VARCHAR(100) NOT NULL COMMENT '活动名称',
    config JSON NOT NULL COMMENT '活动配置',
    start_at BIGINT NOT NULL COMMENT '开始时间',
    end_at BIGINT NOT NULL COMMENT '结束时间',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '活动状态：0-草稿，1-上线',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_category (category),
    INDEX idx_status_time (status, start_at, end_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='活动表';

-- 用户参与记录表
CREATE TABLE IF NOT EXISTS activity_participations (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    activity_id BIGINT NOT NULL COMMENT '活动ID',
    user_id VARCHAR(50) NOT NULL COMMENT '用户ID',
    game_type VARCHAR(50) NOT NULL COMMENT '玩法类型',
    game_target VARCHAR(50) NOT NULL COMMENT '具体玩法标识',
    state VARCHAR(20) NOT NULL COMMENT '参与状态',
    extra JSON COMMENT '额外数据',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_activity_user (activity_id, user_id),
    INDEX idx_user_state (user_id, state)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户参与记录表';

-- 奖品发放记录表
CREATE TABLE IF NOT EXISTS prize_records (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    activity_id BIGINT NOT NULL COMMENT '活动ID',
    user_id VARCHAR(50) NOT NULL COMMENT '用户ID',
    prize_type VARCHAR(50) NOT NULL COMMENT '奖品类型',
    prize_id VARCHAR(50) NOT NULL COMMENT '奖品ID',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '发放状态：0-待发放，1-已发放，2-发放失败',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_activity_user (activity_id, user_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='奖品发放记录表'; 