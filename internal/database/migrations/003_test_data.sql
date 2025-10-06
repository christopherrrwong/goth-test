-- Insert test data into sso_integration_mapping
INSERT INTO `sso_integration_mapping` (`ssousername`, `aclusername`) VALUES
('john.doe@example.com', 'john.doe'),
('jane.smith@example.com', 'jane.smith'),
('test@gmail.com', 'test');


-- Note: For api_token and acl_qr, we'll let the application generate the tokens
-- But here's an example of how the data would look after the application runs:

-- Example of what api_token might contain after app generates tokens
-- INSERT INTO `api_token` (`username`, `token`) VALUES
-- ('john.doe', '1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef'),
-- ('jane.smith', 'abcdef1234567890abcdef1234567890abcdef1234567890abcdef12345678');

-- Example of what acl_qr might contain after app generates tokens
-- INSERT INTO `acl_qr` (`uuid`, `token`, `device_name`) VALUES
-- ('uuid-1', '1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef', 'uuid-1'),
-- ('uuid-2', 'abcdef1234567890abcdef1234567890abcdef1234567890abcdef12345678', 'uuid-2');
