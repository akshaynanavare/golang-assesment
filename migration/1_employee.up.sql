-- Create employees table
CREATE TABLE IF NOT EXISTS employees (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    position TEXT NOT NULL,
    salary BIGINT
);

-- Insert dummy data
INSERT INTO employees (id, name, position, salary) VALUES
    ('3cbe7dd8-1a07-4e02-af1b-92c8f6b5854a', 'John Doe', 'Manager', 50000),
    ('37a1807e-b97d-47f1-a017-2d5d9b15b42a', 'Jane Smith', 'Engineer', 60000),
    ('c301dcb5-ff08-47e2-8d4a-b5d825f3188f', 'Alice Johnson', 'Developer', 70000),
    ('d80cf4f7-3024-4312-a33d-10620e396d72', 'Bob Brown', 'Designer', 55000),
    ('d5d23b26-83c4-4f94-82f8-dedaa25e90c4', 'Eve Wilson', 'Accountant', 65000),
    ('d31a8fa1-7e3d-4c25-84ec-d3acdb0d2892', 'Michael Lee', 'HR Manager', 48000),
    ('5e32a7c3-9716-45b0-aee3-0c1d03291d3b', 'Emily Davis', 'Sales Executive', 75000),
    ('74f6c018-fa35-42b1-af9e-4c3e52540563', 'David Clark', 'Marketing Manager', 62000),
    ('f7483981-1f06-4b2b-82f2-d71d4b5e494e', 'Sarah White', 'Customer Support', 56000),
    ('c75ac497-1d69-4e6e-aa4b-15f93b1bc2a0', 'Robert Harris', 'Software Tester', 58000),
    ('12898517-7d44-4564-9c9c-3b8db82e059d', 'Laura Johnson', 'Manager', 55000),
    ('2f80f5b2-b301-4c5c-b3c5-7c80f5c8c3f1', 'Chris Brown', 'Engineer', 63000),
    ('ea2ec65c-fd2e-4a1f-968d-08a3c6b87e2d', 'Michael Smith', 'Developer', 71000),
    ('ee6e7a60-21fc-4f84-8134-35bb9c184d70', 'Emma Davis', 'Designer', 58000),
    ('65a21aae-7385-4fe5-ba39-5a45a6a4998b', 'Matthew Wilson', 'Accountant', 67000),
    ('950a3c48-07d2-48df-b379-107f6b9871e5', 'Amanda Lee', 'HR Manager', 49000),
    ('a3fbae8a-ff4c-4399-844b-b7d8ed6e0f7f', 'Olivia Davis', 'Sales Executive', 73000),
    ('d9185b62-8b96-43df-93f5-9f8a22504df3', 'Daniel Clark', 'Marketing Manager', 64000),
    ('5f4ae462-9b77-413b-9e51-5710fc52059a', 'James White', 'Customer Support', 58000),
    ('b83b27f3-7461-4de5-b5b2-90cf37c7a2dc', 'Sarah Harris', 'Software Tester', 60000);
