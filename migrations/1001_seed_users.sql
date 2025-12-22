-- =========================
-- SEED USERS
-- =========================

INSERT INTO users (name, email, password, is_active)
VALUES
(
    'Admin',
    'admin@demo.local',
    '$2a$10$CwTycUXWue0Thq9StjUM0uJ8b5bYV3G2M0xwXK5Y3nRk0Wc2q0N1K',
    TRUE
),
(
    'User',
    'user@demo.local',
    '$2a$10$CwTycUXWue0Thq9StjUM0uJ8b5bYV3G2M0xwXK5Y3nRk0Wc2q0N1K',
    TRUE
)
ON DUPLICATE KEY UPDATE
    name = VALUES(name),
    is_active = VALUES(is_active);

-- =========================
-- ASSIGN ROLES TO USERS
-- =========================

-- ADMIN USER → admin role
INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id
FROM users u
JOIN roles r ON r.name = 'admin'
WHERE u.email = 'admin@demo.local'
ON DUPLICATE KEY UPDATE
    user_id = user_id;

-- NORMAL USER → user role
INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id
FROM users u
JOIN roles r ON r.name = 'user'
WHERE u.email = 'user@demo.local'
ON DUPLICATE KEY UPDATE
    user_id = user_id;
