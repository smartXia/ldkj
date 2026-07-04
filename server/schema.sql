-- MySQL schema for the official website and CMS.
-- Keep existing table and column names used by the Go backend compatible.

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE IF NOT EXISTS site_config (
  id BIGINT PRIMARY KEY COMMENT '站点配置 ID',
  site_title VARCHAR(255) NOT NULL DEFAULT '' COMMENT '站点标题',
  logo_url VARCHAR(500) NOT NULL DEFAULT '' COMMENT '站点 Logo 图片地址',
  contact TEXT NOT NULL COMMENT '联系信息',
  copyright VARCHAR(255) NOT NULL DEFAULT '' COMMENT '版权声明',
  footer_links TEXT NOT NULL COMMENT '页脚链接配置，建议存 JSON 字符串',
  icp_no VARCHAR(100) NOT NULL DEFAULT '' COMMENT 'ICP备案号',
  public_security_no VARCHAR(100) NOT NULL DEFAULT '' COMMENT '公安备案号',
  analytics_code MEDIUMTEXT NULL COMMENT '第三方统计代码',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='站点基础配置表';

CREATE TABLE IF NOT EXISTS services (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '服务 ID',
  title VARCHAR(255) NOT NULL COMMENT '服务名称',
  slug VARCHAR(160) NOT NULL COMMENT '服务访问标识',
  subtitle VARCHAR(500) NOT NULL DEFAULT '' COMMENT '服务副标题',
  summary VARCHAR(500) NOT NULL DEFAULT '' COMMENT '服务摘要',
  cover_url VARCHAR(500) NOT NULL DEFAULT '' COMMENT '服务封面图',
  icon_url VARCHAR(500) NOT NULL DEFAULT '' COMMENT '服务图标',
  content MEDIUMTEXT NOT NULL COMMENT '服务详情内容',
  highlights TEXT NOT NULL COMMENT '服务亮点，建议存 JSON 字符串',
  process_steps TEXT NOT NULL COMMENT '服务流程，建议存 JSON 字符串',
  sort_order INT NOT NULL DEFAULT 0 COMMENT '排序值',
  status ENUM('draft', 'published', 'archived') NOT NULL DEFAULT 'draft' COMMENT '发布状态',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  UNIQUE KEY uk_services_slug (slug),
  KEY idx_services_status_sort (status, sort_order)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='服务模块表';

CREATE TABLE IF NOT EXISTS banners (
  id BIGINT PRIMARY KEY COMMENT '轮播图 ID',
  title VARCHAR(255) NOT NULL DEFAULT '' COMMENT '轮播图标题',
  subtitle VARCHAR(500) NOT NULL DEFAULT '' COMMENT '轮播图副标题',
  image_url VARCHAR(500) NOT NULL DEFAULT '' COMMENT '轮播图图片地址',
  link_url VARCHAR(500) NOT NULL DEFAULT '' COMMENT '跳转链接地址',
  button_text VARCHAR(100) NOT NULL DEFAULT '' COMMENT '按钮文案',
  position_key VARCHAR(100) NOT NULL DEFAULT 'home_hero' COMMENT '展示位置标识',
  sort_order INT NOT NULL DEFAULT 0 COMMENT '排序值',
  is_published BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否发布',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  KEY idx_banners_position (position_key, is_published, sort_order)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='首页轮播图表';

CREATE TABLE IF NOT EXISTS seo_settings (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'SEO 配置 ID',
  page_key VARCHAR(100) NOT NULL COMMENT '页面唯一标识',
  title VARCHAR(255) NOT NULL DEFAULT '' COMMENT '页面标题',
  description VARCHAR(500) NOT NULL DEFAULT '' COMMENT '页面描述',
  keywords VARCHAR(500) NOT NULL DEFAULT '' COMMENT '页面关键词',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  UNIQUE KEY uk_seo_page_key (page_key)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='页面 SEO 配置表';

CREATE TABLE IF NOT EXISTS cases (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '案例 ID',
  service_id BIGINT NULL COMMENT '关联服务 ID',
  title VARCHAR(255) NOT NULL COMMENT '案例标题',
  slug VARCHAR(160) NOT NULL COMMENT '案例访问标识',
  industry VARCHAR(100) NOT NULL DEFAULT '' COMMENT '所属行业',
  platform VARCHAR(100) NOT NULL DEFAULT '' COMMENT '投放平台',
  strategy VARCHAR(100) NOT NULL DEFAULT '' COMMENT '营销打法或赛道',
  cover_url VARCHAR(500) NOT NULL DEFAULT '' COMMENT '案例封面图地址',
  summary VARCHAR(500) NOT NULL DEFAULT '' COMMENT '案例摘要',
  content MEDIUMTEXT NOT NULL COMMENT '案例详情内容',
  core_metrics TEXT NOT NULL COMMENT '核心指标数据，建议存 JSON 字符串',
  sort_order INT NOT NULL DEFAULT 0 COMMENT '排序值',
  status ENUM('draft', 'published', 'archived') NOT NULL DEFAULT 'draft' COMMENT '发布状态',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  UNIQUE KEY uk_cases_slug (slug),
  KEY idx_cases_status (status),
  KEY idx_cases_filters (industry, platform, strategy),
  KEY idx_cases_service (service_id),
  CONSTRAINT fk_cases_service FOREIGN KEY (service_id) REFERENCES services(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='客户案例表';

CREATE TABLE IF NOT EXISTS news (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '资讯 ID',
  title VARCHAR(255) NOT NULL COMMENT '资讯标题',
  slug VARCHAR(160) NOT NULL COMMENT '资讯访问标识',
  category VARCHAR(100) NOT NULL DEFAULT '' COMMENT '资讯分类',
  author VARCHAR(100) NOT NULL DEFAULT '' COMMENT '作者',
  cover_url VARCHAR(500) NOT NULL DEFAULT '' COMMENT '资讯封面图地址',
  summary VARCHAR(500) NOT NULL DEFAULT '' COMMENT '资讯摘要',
  content MEDIUMTEXT NOT NULL COMMENT '资讯正文内容',
  status ENUM('draft', 'published', 'archived') NOT NULL DEFAULT 'draft' COMMENT '发布状态',
  published_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  UNIQUE KEY uk_news_slug (slug),
  KEY idx_news_status_category (status, category),
  KEY idx_news_published_at (published_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='新闻资讯表';

CREATE TABLE IF NOT EXISTS static_pages (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '静态页面 ID',
  page_key VARCHAR(100) NOT NULL COMMENT '页面标识，例如 about、partner',
  title VARCHAR(255) NOT NULL DEFAULT '' COMMENT '页面标题',
  content MEDIUMTEXT NOT NULL COMMENT '富文本内容',
  extra_data MEDIUMTEXT NULL COMMENT '扩展配置，建议存 JSON 字符串',
  status ENUM('draft', 'published', 'archived') NOT NULL DEFAULT 'draft' COMMENT '发布状态',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  UNIQUE KEY uk_static_pages_page_key (page_key),
  KEY idx_static_pages_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='静态页面内容表';

CREATE TABLE IF NOT EXISTS partner_faqs (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'FAQ ID',
  question VARCHAR(255) NOT NULL COMMENT '问题',
  answer TEXT NOT NULL COMMENT '答案',
  sort_order INT NOT NULL DEFAULT 0 COMMENT '排序值',
  is_published BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否发布',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  KEY idx_partner_faqs_publish_sort (is_published, sort_order)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='合作页 FAQ 表';

CREATE TABLE IF NOT EXISTS lead_forms (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '线索表单 ID',
  form_type VARCHAR(50) NOT NULL DEFAULT 'contact' COMMENT '表单类型：home、partner 等',
  name VARCHAR(100) NOT NULL COMMENT '联系人姓名',
  phone VARCHAR(50) NOT NULL COMMENT '联系电话',
  company VARCHAR(255) NOT NULL COMMENT '公司名称',
  position VARCHAR(100) NOT NULL DEFAULT '' COMMENT '职位',
  email VARCHAR(255) NOT NULL DEFAULT '' COMMENT '邮箱地址',
  requirement TEXT NOT NULL COMMENT '需求描述',
  interest VARCHAR(500) NOT NULL DEFAULT '' COMMENT '意向服务或资料',
  source VARCHAR(255) NOT NULL DEFAULT '' COMMENT '线索来源页面',
  status ENUM('new', 'processed') NOT NULL DEFAULT 'new' COMMENT '处理状态',
  processed_at DATETIME NULL COMMENT '处理时间',
  processed_by BIGINT NULL COMMENT '处理人 ID',
  ip_address VARCHAR(45) NOT NULL DEFAULT '' COMMENT '提交 IP',
  user_agent VARCHAR(500) NOT NULL DEFAULT '' COMMENT '浏览器 User-Agent',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  KEY idx_lead_forms_status (status),
  KEY idx_lead_forms_created_at (created_at),
  KEY idx_lead_forms_source (source)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='线索表单表';

CREATE TABLE IF NOT EXISTS media_assets (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '媒体资源 ID',
  original_name VARCHAR(255) NOT NULL COMMENT '原始文件名',
  file_name VARCHAR(255) NOT NULL COMMENT '存储文件名',
  file_path VARCHAR(500) NOT NULL COMMENT '本地存储路径',
  file_url VARCHAR(500) NOT NULL COMMENT '可访问 URL',
  mime_type VARCHAR(100) NOT NULL DEFAULT '' COMMENT 'MIME 类型',
  file_size BIGINT NOT NULL DEFAULT 0 COMMENT '文件大小，单位字节',
  biz_type VARCHAR(100) NOT NULL DEFAULT '' COMMENT '业务类型：banner、case、news、editor 等',
  ref_table VARCHAR(100) NOT NULL DEFAULT '' COMMENT '关联业务表',
  ref_id BIGINT NULL COMMENT '关联业务 ID',
  uploader_id BIGINT NULL COMMENT '上传人 ID',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  KEY idx_media_assets_biz (biz_type, ref_table, ref_id),
  KEY idx_media_assets_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='本地文件上传资源表';

CREATE TABLE IF NOT EXISTS admin_users (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '管理员 ID',
  username VARCHAR(100) NOT NULL COMMENT '登录账号',
  password_hash VARCHAR(255) NOT NULL COMMENT '密码哈希',
  real_name VARCHAR(100) NOT NULL DEFAULT '' COMMENT '真实姓名',
  email VARCHAR(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  phone VARCHAR(50) NOT NULL DEFAULT '' COMMENT '手机',
  status ENUM('active', 'disabled') NOT NULL DEFAULT 'active' COMMENT '账号状态',
  last_login_at DATETIME NULL COMMENT '最后登录时间',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  UNIQUE KEY uk_admin_users_username (username),
  KEY idx_admin_users_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='后台管理员表';

CREATE TABLE IF NOT EXISTS roles (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '角色 ID',
  code VARCHAR(100) NOT NULL COMMENT '角色编码',
  name VARCHAR(100) NOT NULL COMMENT '角色名称',
  description VARCHAR(500) NOT NULL DEFAULT '' COMMENT '角色说明',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  UNIQUE KEY uk_roles_code (code)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

CREATE TABLE IF NOT EXISTS permissions (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '权限 ID',
  code VARCHAR(100) NOT NULL COMMENT '权限编码',
  name VARCHAR(100) NOT NULL COMMENT '权限名称',
  module VARCHAR(100) NOT NULL DEFAULT '' COMMENT '所属模块',
  description VARCHAR(500) NOT NULL DEFAULT '' COMMENT '权限说明',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  UNIQUE KEY uk_permissions_code (code),
  KEY idx_permissions_module (module)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限表';

CREATE TABLE IF NOT EXISTS admin_user_roles (
  user_id BIGINT NOT NULL COMMENT '管理员 ID',
  role_id BIGINT NOT NULL COMMENT '角色 ID',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (user_id, role_id),
  CONSTRAINT fk_admin_user_roles_user FOREIGN KEY (user_id) REFERENCES admin_users(id) ON DELETE CASCADE,
  CONSTRAINT fk_admin_user_roles_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员角色关联表';

CREATE TABLE IF NOT EXISTS role_permissions (
  role_id BIGINT NOT NULL COMMENT '角色 ID',
  permission_id BIGINT NOT NULL COMMENT '权限 ID',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (role_id, permission_id),
  CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
  CONSTRAINT fk_role_permissions_permission FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色权限关联表';

CREATE TABLE IF NOT EXISTS operation_logs (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '操作日志 ID',
  user_id BIGINT NULL COMMENT '操作人 ID',
  username VARCHAR(100) NOT NULL DEFAULT '' COMMENT '操作账号',
  action VARCHAR(100) NOT NULL COMMENT '操作类型：login、create、update、delete 等',
  module VARCHAR(100) NOT NULL DEFAULT '' COMMENT '业务模块',
  target_table VARCHAR(100) NOT NULL DEFAULT '' COMMENT '目标表',
  target_id BIGINT NULL COMMENT '目标记录 ID',
  detail MEDIUMTEXT NULL COMMENT '操作详情，建议存 JSON 字符串',
  ip_address VARCHAR(45) NOT NULL DEFAULT '' COMMENT '操作 IP',
  user_agent VARCHAR(500) NOT NULL DEFAULT '' COMMENT '浏览器 User-Agent',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  KEY idx_operation_logs_user (user_id),
  KEY idx_operation_logs_action (action),
  KEY idx_operation_logs_created_at (created_at),
  CONSTRAINT fk_operation_logs_user FOREIGN KEY (user_id) REFERENCES admin_users(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='后台操作日志表';

CREATE TABLE IF NOT EXISTS email_settings (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '邮件配置 ID',
  smtp_host VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'SMTP 主机',
  smtp_port INT NOT NULL DEFAULT 465 COMMENT 'SMTP 端口',
  smtp_user VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'SMTP 用户名',
  smtp_password VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'SMTP 密码或授权码',
  sender_email VARCHAR(255) NOT NULL DEFAULT '' COMMENT '发件人邮箱',
  sender_name VARCHAR(100) NOT NULL DEFAULT '' COMMENT '发件人名称',
  recipients TEXT NOT NULL COMMENT '通知收件人，多个邮箱可用逗号分隔或存 JSON 字符串',
  is_ssl BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否启用 SSL',
  is_enabled BOOLEAN NOT NULL DEFAULT FALSE COMMENT '是否启用通知',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='邮件通知配置表';

CREATE TABLE IF NOT EXISTS email_notifications (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '邮件通知记录 ID',
  lead_form_id BIGINT NULL COMMENT '关联线索表单 ID',
  recipient VARCHAR(255) NOT NULL DEFAULT '' COMMENT '收件人',
  subject VARCHAR(255) NOT NULL DEFAULT '' COMMENT '邮件主题',
  content MEDIUMTEXT NOT NULL COMMENT '邮件内容',
  status ENUM('pending', 'sent', 'failed') NOT NULL DEFAULT 'pending' COMMENT '发送状态',
  error_message TEXT NULL COMMENT '失败原因',
  sent_at DATETIME NULL COMMENT '发送时间',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  KEY idx_email_notifications_form (lead_form_id),
  KEY idx_email_notifications_status (status),
  CONSTRAINT fk_email_notifications_form FOREIGN KEY (lead_form_id) REFERENCES lead_forms(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='邮件通知发送记录表';

CREATE TABLE IF NOT EXISTS analytics_settings (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '统计配置 ID',
  provider VARCHAR(100) NOT NULL DEFAULT '' COMMENT '统计工具名称',
  tracking_id VARCHAR(255) NOT NULL DEFAULT '' COMMENT '跟踪 ID',
  script_code MEDIUMTEXT NULL COMMENT '统计脚本代码',
  is_enabled BOOLEAN NOT NULL DEFAULT FALSE COMMENT '是否启用',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='第三方数据统计配置表';

CREATE TABLE IF NOT EXISTS system_settings (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '系统配置 ID',
  setting_key VARCHAR(100) NOT NULL COMMENT '配置键',
  setting_value MEDIUMTEXT NULL COMMENT '配置值',
  description VARCHAR(500) NOT NULL DEFAULT '' COMMENT '配置说明',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  UNIQUE KEY uk_system_settings_key (setting_key)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统配置扩展表';

SET FOREIGN_KEY_CHECKS = 1;

INSERT INTO site_config (id, site_title, logo_url, contact, copyright, footer_links, icp_no, public_security_no)
VALUES (1, '南京灵动信息技术有限公司', '', '', '© 南京灵动信息技术有限公司', '[]', '', '')
ON DUPLICATE KEY UPDATE id = id;

INSERT INTO banners (id, title, subtitle, image_url, link_url, button_text, position_key, sort_order, is_published)
VALUES (1, '南京灵动信息技术有限公司', '数字营销与内容增长服务', '', '', '立即咨询', 'home_hero', 0, TRUE)
ON DUPLICATE KEY UPDATE id = id;

INSERT INTO seo_settings (page_key, title, description, keywords)
VALUES
  ('home', '南京灵动信息技术有限公司', '南京灵动信息技术有限公司官方网站', '南京灵动,数字营销,内容增长'),
  ('services', '服务模块 - 南京灵动信息技术有限公司', '了解南京灵动信息技术有限公司提供的数字营销与内容增长服务', '服务模块,数字营销,内容增长'),
  ('cases', '服务案例 - 南京灵动信息技术有限公司', '查看南京灵动信息技术有限公司服务案例', '服务案例,营销案例'),
  ('news', '营销资讯 - 南京灵动信息技术有限公司', '获取灵动动态、营销洞察和行业资讯', '营销资讯,灵动动态,行业资讯'),
  ('about', '关于我们 - 南京灵动信息技术有限公司', '了解南京灵动信息技术有限公司', '关于我们,公司介绍'),
  ('partner', '合作咨询 - 南京灵动信息技术有限公司', '提交合作咨询表单，与南京灵动建立联系', '合作咨询,商务合作')
ON DUPLICATE KEY UPDATE page_key = VALUES(page_key);

INSERT INTO roles (code, name, description)
VALUES
  ('super_admin', '超级管理员', '拥有全部后台管理权限'),
  ('editor', '编辑', '负责内容录入和编辑'),
  ('operator', '运营', '负责表单查看、内容运营等操作')
ON DUPLICATE KEY UPDATE name = VALUES(name), description = VALUES(description);

INSERT INTO permissions (code, name, module, description)
VALUES
  ('content:read', '内容查看', 'content', '查看案例、资讯、服务和静态页内容'),
  ('content:write', '内容编辑', 'content', '新增和编辑案例、资讯、服务和静态页内容'),
  ('content:delete', '内容删除', 'content', '删除内容记录'),
  ('banner:manage', 'Banner 管理', 'banner', '管理首页 Banner'),
  ('form:read', '表单查看', 'form', '查看线索表单'),
  ('form:process', '表单处理', 'form', '标记线索处理状态和导出数据'),
  ('settings:manage', '基础设置', 'settings', '管理站点配置、SEO、邮件、统计等设置'),
  ('user:manage', '用户权限管理', 'user', '管理后台用户、角色和权限'),
  ('log:read', '日志查看', 'log', '查看后台操作日志')
ON DUPLICATE KEY UPDATE name = VALUES(name), module = VALUES(module), description = VALUES(description);

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON r.code = 'super_admin'
ON DUPLICATE KEY UPDATE role_id = role_id;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON r.code = 'editor' AND p.code IN ('content:read', 'content:write', 'banner:manage')
ON DUPLICATE KEY UPDATE role_id = role_id;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON r.code = 'operator' AND p.code IN ('content:read', 'form:read', 'form:process', 'log:read')
ON DUPLICATE KEY UPDATE role_id = role_id;

INSERT INTO email_settings (id, smtp_host, smtp_port, smtp_user, smtp_password, sender_email, sender_name, recipients, is_ssl, is_enabled)
VALUES (1, '', 465, '', '', '', '南京灵动信息技术有限公司', '', TRUE, FALSE)
ON DUPLICATE KEY UPDATE id = id;

INSERT INTO analytics_settings (id, provider, tracking_id, script_code, is_enabled)
VALUES (1, '', '', '', FALSE)
ON DUPLICATE KEY UPDATE id = id;

INSERT INTO system_settings (setting_key, setting_value, description)
VALUES
  ('upload_root', '/oss', '本地文件上传根目录'),
  ('site_domain', 'www.ilingdong.cn', '网站域名')
ON DUPLICATE KEY UPDATE setting_value = VALUES(setting_value), description = VALUES(description);
