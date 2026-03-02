-- Docker集群管理模块初始化 SQL
-- 执行此SQL前，请确保已经运行过后端程序（表已创建）

-- ============================================
-- 1. 插入菜单数据
-- ============================================

-- 检查是否已存在，避免重复插入
SET @parent_id = 0;

-- 插入父菜单（Docker集群管理）
INSERT INTO `sys_base_menus` (
    `created_at`, `updated_at`, `deleted_at`,
    `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`,
    `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`, `transition_type`, `active_name`
) VALUES (
    NOW(), NOW(), NULL,
    0, 0, 'dockerCluster', 'dockerCluster', 0, 'view/dockerCluster/dockerCluster.vue', 0,
    0, 0, 'Docker集群管理', 'cpu', 0, '', ''
);

-- 获取刚插入的菜单ID
SET @parent_id = LAST_INSERT_ID();

-- ============================================
-- 2. 插入按钮权限数据
-- ============================================

-- 新增按钮
INSERT INTO `sys_base_menu_btns` (`created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES
(NOW(), NOW(), NULL, '新增', 'dockerCluster_btn_add', @parent_id);

-- 删除按钮
INSERT INTO `sys_base_menu_btns` (`created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES
(NOW(), NOW(), NULL, '删除', 'dockerCluster_btn_delete', @parent_id);

-- 编辑按钮
INSERT INTO `sys_base_menu_btns` (`created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES
(NOW(), NOW(), NULL, '编辑', 'dockerCluster_btn_edit', @parent_id);

-- 查看凭证按钮
INSERT INTO `sys_base_menu_btns` (`created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `sys_base_menu_id`) VALUES
(NOW(), NOW(), NULL, '查看凭证', 'dockerCluster_btn_credentials', @parent_id);

-- ============================================
-- 3. 插入API接口权限数据
-- ============================================

-- 创建Docker集群
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES
(NOW(), NOW(), NULL, '/dockerCluster/createDockerCluster', '创建Docker集群', 'Docker集群管理', 'POST');

-- 删除Docker集群
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES
(NOW(), NOW(), NULL, '/dockerCluster/deleteDockerCluster', '删除Docker集群', 'Docker集群管理', 'DELETE');

-- 批量删除Docker集群
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES
(NOW(), NOW(), NULL, '/dockerCluster/deleteDockerClusterByIds', '批量删除Docker集群', 'Docker集群管理', 'DELETE');

-- 更新Docker集群
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES
(NOW(), NOW(), NULL, '/dockerCluster/updateDockerCluster', '更新Docker集群', 'Docker集群管理', 'PUT');

-- 根据ID查询Docker集群
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES
(NOW(), NOW(), NULL, '/dockerCluster/findDockerCluster', '根据ID查询Docker集群', 'Docker集群管理', 'GET');

-- 获取Docker集群列表
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES
(NOW(), NOW(), NULL, '/dockerCluster/getDockerClusterList', '获取Docker集群列表', 'Docker集群管理', 'GET');

-- 获取Docker集群凭证
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES
(NOW(), NOW(), NULL, '/dockerCluster/getDockerClusterCredentials', '获取Docker集群凭证', 'Docker集群管理', 'GET');

-- 获取所有Docker集群
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES
(NOW(), NOW(), NULL, '/dockerCluster/getAllDockerClusters', '获取所有Docker集群', 'Docker集群管理', 'GET');

-- ============================================
-- 4. 将菜单授权给超级管理员角色（authorityId = 888）
-- ============================================

-- 插入菜单-角色关联
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES
(888, @parent_id);

-- 获取按钮ID并授权给超级管理员
-- 新增按钮
SET @btn_add_id = (SELECT id FROM `sys_base_menu_btns` WHERE `sys_base_menu_id` = @parent_id AND `desc` = 'dockerCluster_btn_add' LIMIT 1);
INSERT INTO `sys_authority_btns` (`authority_id`, `menu_id`, `sys_base_menu_btn_id`) VALUES
(888, @parent_id, @btn_add_id);

-- 删除按钮
SET @btn_delete_id = (SELECT id FROM `sys_base_menu_btns` WHERE `sys_base_menu_id` = @parent_id AND `desc` = 'dockerCluster_btn_delete' LIMIT 1);
INSERT INTO `sys_authority_btns` (`authority_id`, `menu_id`, `sys_base_menu_btn_id`) VALUES
(888, @parent_id, @btn_delete_id);

-- 编辑按钮
SET @btn_edit_id = (SELECT id FROM `sys_base_menu_btns` WHERE `sys_base_menu_id` = @parent_id AND `desc` = 'dockerCluster_btn_edit' LIMIT 1);
INSERT INTO `sys_authority_btns` (`authority_id`, `menu_id`, `sys_base_menu_btn_id`) VALUES
(888, @parent_id, @btn_edit_id);

-- 查看凭证按钮
SET @btn_credentials_id = (SELECT id FROM `sys_base_menu_btns` WHERE `sys_base_menu_id` = @parent_id AND `desc` = 'dockerCluster_btn_credentials' LIMIT 1);
INSERT INTO `sys_authority_btns` (`authority_id`, `menu_id`, `sys_base_menu_btn_id`) VALUES
(888, @parent_id, @btn_credentials_id);

-- ============================================
-- 5. 插入casbin权限规则（允许超级管理员访问）
-- ============================================

-- POST /dockerCluster/createDockerCluster
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
('p', '888', '/dockerCluster/createDockerCluster', 'POST', '', '', '');

-- DELETE /dockerCluster/deleteDockerCluster
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
('p', '888', '/dockerCluster/deleteDockerCluster', 'DELETE', '', '', '');

-- DELETE /dockerCluster/deleteDockerClusterByIds
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
('p', '888', '/dockerCluster/deleteDockerClusterByIds', 'DELETE', '', '', '');

-- PUT /dockerCluster/updateDockerCluster
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
('p', '888', '/dockerCluster/updateDockerCluster', 'PUT', '', '', '');

-- GET /dockerCluster/findDockerCluster
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
('p', '888', '/dockerCluster/findDockerCluster', 'GET', '', '', '');

-- GET /dockerCluster/getDockerClusterList
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
('p', '888', '/dockerCluster/getDockerClusterList', 'GET', '', '', '');

-- GET /dockerCluster/getDockerClusterCredentials
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
('p', '888', '/dockerCluster/getDockerClusterCredentials', 'GET', '', '', '');

-- GET /dockerCluster/getAllDockerClusters
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
('p', '888', '/dockerCluster/getAllDockerClusters', 'GET', '', '', '');

-- ============================================
-- 初始化完成
-- ============================================
SELECT 'Docker集群管理模块初始化完成！' AS result;
