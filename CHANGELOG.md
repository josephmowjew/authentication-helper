# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.1] - 2024-02-18
### Changed
- Updated package import path from 'josephmowjew' to 'josephmojoo'

### Security
- Replaced hardcoded credentials with placeholder values in examples

## [0.1.0] - 2025-02-18
### Added
- Initial release of the authentication helper package
- Authentication client implementation with configurable timeout and base URL
- Password hashing and verification utilities using bcrypt
- Token expiration checking functionality
- Structured response types for authentication data
- Example implementation demonstrating usage

### Security
- Implemented secure password hashing using bcrypt
- Proper error handling for authentication failures
- Timeout configuration for HTTP requests