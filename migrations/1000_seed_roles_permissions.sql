-- ROLES
INSERT INTO roles (name) VALUES
('admin'),
('user')
ON DUPLICATE KEY UPDATE name = name;

-- PERMISSIONS
INSERT INTO permissions (name) VALUES
('user.read'),
('user.create'),
('user.update'),
('user.delete'),
('admin.dashboard')
ON DUPLICATE KEY UPDATE name = name;

-- ADMIN ROLE → ALL PERMISSIONS
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r, permissions p
WHERE r.name = 'admin'
ON DUPLICATE KEY UPDATE role_id = role_id;

-- USER ROLE → BASIC PERMISSIONS
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON p.name IN ('user.read')
WHERE r.name = 'user'
ON DUPLICATE KEY UPDATE role_id = role_id;
