# Design Document

## Overview

The Telegram Mini App integration enhances the existing quiz application with Telegram-specific features while maintaining full backward compatibility with standard web browsers. The design follows a progressive enhancement approach where Telegram features are layered on top of the existing functionality.

## Architecture

### Core Components

1. **Telegram Detection Service** - Determines if the app is running in Telegram environment
2. **Telegram SDK Manager** - Handles SDK initialization and lifecycle
3. **Player ID Service** - Manages user identification across environments
4. **Telegram UI Components** - Handles Telegram-specific UI elements like back button
5. **API Client Enhancement** - Extends existing API client with Telegram user context

### Integration Points

- **Global App Initialization** - SDK setup in app.html
- **API Client** - Enhanced to include Telegram user ID in headers
- **Game Page** - Back button integration
- **Player ID Management** - Unified across Telegram and web environments

## Components and Interfaces

### Telegram Detection Service

```typescript
interface TelegramDetectionService {
  isTelegramEnvironment(): boolean;
  getTelegramWebApp(): WebApp | null;
  isInitialized(): boolean;
}
```

**Location:** `src/lib/telegram/detection.ts`

**Responsibilities:**
- Detect Telegram Mini App environment
- Provide access to Telegram WebApp instance
- Handle initialization state

### Telegram SDK Manager

```typescript
interface TelegramSDKManager {
  initialize(): Promise<void>;
  isReady(): boolean;
  getWebApp(): WebApp | null;
  cleanup(): void;
}
```

**Location:** `src/lib/telegram/sdk.ts`

**Responsibilities:**
- Initialize @telegram-apps/sdk
- Manage SDK lifecycle
- Provide centralized access to Telegram features
- Handle initialization errors gracefully

### Enhanced Player ID Service

```typescript
interface PlayerIDService {
  getPlayerId(): string;
  getTelegramUserId(): number | null;
  isUsingTelegramId(): boolean;
}
```

**Location:** `src/lib/services/playerId.ts`

**Responsibilities:**
- Generate or retrieve player ID based on environment
- Extract Telegram user ID when available
- Maintain backward compatibility with existing ID logic
- Handle fallback scenarios

### Telegram UI Components

```typescript
interface TelegramUIManager {
  showBackButton(): void;
  hideBackButton(): void;
  onBackButtonClick(callback: () => void): void;
  isBackButtonVisible(): boolean;
}
```

**Location:** `src/lib/telegram/ui.ts`

**Responsibilities:**
- Manage Telegram back button visibility
- Handle back button interactions
- Provide clean API for UI state management

### Enhanced API Client

**Location:** `src/lib/api/client.ts` (enhanced)

**Changes:**
- Add X-Player-ID header to all requests
- Use Telegram user ID when available
- Fallback to existing player ID logic
- Maintain existing API interface

## Data Models

### Telegram Context

```typescript
interface TelegramContext {
  isInTelegram: boolean;
  userId?: number;
  username?: string;
  firstName?: string;
  lastName?: string;
  languageCode?: string;
}
```

### Player Identification

```typescript
interface PlayerIdentification {
  playerId: string;
  source: 'telegram' | 'generated';
  telegramUserId?: number;
}
```

## Error Handling

### SDK Initialization Failures
- Log errors to console
- Continue with standard web functionality
- Provide fallback for all Telegram-dependent features

### Telegram API Failures
- Graceful degradation to web-only features
- Retry mechanisms for transient failures
- User-friendly error messages when appropriate

### Player ID Fallbacks
- Always ensure a valid player ID is available
- Seamless transition between Telegram and generated IDs
- Maintain session consistency across failures

## Testing Strategy

### Unit Tests
- Telegram detection logic
- Player ID service with various scenarios
- SDK manager initialization and error handling
- UI component state management

### Integration Tests
- API client with different player ID sources
- End-to-end game flow in both environments
- Back button navigation behavior

### Environment Testing
- Test in actual Telegram Mini App environment
- Verify fallback behavior in standard browsers
- Cross-browser compatibility testing

## Implementation Phases

### Phase 1: Foundation
- Install and configure @telegram-apps/sdk
- Implement detection service
- Create basic SDK manager

### Phase 2: Player Identification
- Enhance player ID service
- Update API client with header injection
- Implement fallback mechanisms

### Phase 3: UI Integration
- Add back button functionality to game page
- Implement UI state management
- Handle navigation logic

### Phase 4: Testing & Polish
- Comprehensive testing across environments
- Error handling refinement
- Performance optimization

## Security Considerations

- Validate Telegram user data before using
- Sanitize user inputs from Telegram context
- Implement proper error boundaries
- Avoid exposing sensitive information in logs

## Performance Considerations

- Lazy load Telegram SDK when not needed
- Minimize impact on standard web performance
- Efficient detection mechanisms
- Proper cleanup of Telegram resources