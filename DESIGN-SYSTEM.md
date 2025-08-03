# **Design System (v3)**

## **Overview**

The application design is **playful, vibrant, and modern**. Its implementation fully relies on **daisyUI components** with Tailwind CSS v4 for maximum consistency and development speed. A global width constraint (max-w-md) gives the application a focused "mobile-native" feel across all devices, optimized for Telegram Mini App usage.

## **CSS Setup**

This project uses **Tailwind CSS v4** with daisyUI plugin in zero-config mode. All customization is handled directly in the app.css file using the new `@plugin` syntax. No separate tailwind.config.js file is used.

**Key Features:**
- Custom daisyUI theme with CSS custom properties
- Pribambas font integration via `@font-face`
- Component-specific utility classes
- Telegram Mini App optimizations

References: _/web/src/app.css_

## **Dependencies & Versions**

### **Core Styling Dependencies**
- **Tailwind CSS:** `@tailwindcss/cli@^4.0.0`, `@tailwindcss/postcss@^4.1.11`, `@tailwindcss/vite@^4.0.0`
- **daisyUI:** `^5.0.46`

### **Framework & Build Tools**
- **SvelteKit:** `^2.5.27`
- **Svelte:** `^5.0.0`
- **Vite:** `^5.4.4`
- **TypeScript:** `^5.5.0`

### **Telegram Integration**
- **Telegram Apps SDK:** `@telegram-apps/sdk@^3.11.4`

### **Additional Libraries**
- **Canvas Confetti:** `canvas-confetti@^1.9.3` (for celebration effects)

**Note:** The project uses Tailwind CSS v4. The zero-config approach with `@plugin` syntax in CSS files is a new feature that eliminates the need for traditional `tailwind.config.js` files.

## **Design Tokens**

### **Color Palette**

The application uses a custom daisyUI light theme with carefully selected colors:

| Token | CSS Variable | Usage |
| :---- | :---- | :---- |
| Primary | `--color-primary` | Purple accent, main branding |
| Secondary | `--color-secondary` | Pink buttons, call-to-action |
| Base-100 | `--color-base-100` | White backgrounds |
| Base-content | `--color-base-content` | Dark text on light backgrounds |
| Primary-content | `--color-primary-content` | White text on primary backgrounds |
| Success | `--color-success` | Correct answers, positive feedback |
| Error | `--color-error` | Wrong answers, error states |

### **Typography**

**Primary Font:** Pribambas (loaded via `@font-face` from `/fonts/pribambas.ttf`)

| Element | Classes | Usage |
| :---- | :---- | :---- |
| App Title | `text-5xl text-main-font` | "ИЗИ КВИЗИ!" main heading |
| Game Title | `text-5xl font-bold text-main-font` | Quiz game names |
| Question Text | `text-2xl font-semibold` | Quiz questions |
| Button Text | `text-lg` or `text-base` | Interactive elements |
| Body Text | `text-sm` to `text-xl` | General content |

### **Layout System**

* **Container:** `container mx-auto max-w-md p-2` - Global wrapper
* **Card Padding:** `p-4` - Standard internal spacing
* **Content Spacing:** `space-y-3`, `gap-4` - Consistent element spacing
* **Maximum Width:** All main content uses `max-w-md` for mobile-first design

### **Border Radius**

Custom radius values defined in CSS variables:

* **Cards:** `rounded-3xl` (2rem) - Main content containers
* **Buttons/Fields:** `rounded-lg` (0.5rem) - Interactive elements
* **Pills:** `rounded-full` - Progress indicators, badges

## **Component Library**

The application uses a combination of daisyUI components and custom implementations optimized for the quiz experience.

### **Layout Components**

#### **Main Container**
Standard wrapper for all pages:
```svelte
<div class="container mx-auto max-w-md p-2">
  <!-- Page content -->
</div>
```

#### **Card Container**
Primary content wrapper with consistent styling:
```svelte
<div class="card bg-primary rounded-3xl">
  <div class="card-body p-4">
    <!-- Card content -->
  </div>
</div>
```

### **Interactive Components**

#### **Primary Buttons**
Main action buttons with consistent styling:

**Secondary (Pink) - Primary Actions:**
```svelte
<button class="btn btn-secondary btn-lg text-base font-semibold transition-transform">
  Начать игру
</button>
```

**Primary (Purple) - Navigation:**
```svelte
<button class="btn btn-primary transition-transform">
  Далее
</button>
```

**Neutral - Alternative Actions:**
```svelte
<button class="btn btn-neutral btn-lg transition-transform">
  Еще раз
</button>
```

#### **Quiz Answer Options**
Custom button style for quiz answers:
```svelte
<button class="btn-answer-option text-left break-words hyphens-auto">
  1. Answer text here
</button>
```

**With State Classes:**
- `.option-correct` - Green background for correct answers
- `.option-incorrect` - Red background for wrong answers

#### **Input Fields**
Text input with consistent styling:
```svelte
<input 
  type="text" 
  class="input input-lg input-bordered w-full text-sm"
  placeholder="c886247f-9dc6-4e3f-b2f0-50f912438079"
/>
```

### **Feedback Components**

#### **Toast Notifications**
System-wide notifications using custom toast system:
```svelte
<!-- Success -->
<div class="alert alert-success mb-2">
  <span class="text-sm font-medium">Успешно!</span>
</div>

<!-- Error -->
<div class="alert alert-error mb-2">
  <span class="text-sm font-medium">Ошибка!</span>
</div>
```

#### **Answer Feedback**
Immediate feedback after quiz answers:
```svelte
<!-- Correct Answer -->
<div class="mb-4 bg-success text-success-content p-4 rounded-lg">
  <span class="text-xl">Правильно!</span>
</div>

<!-- Wrong Answer -->
<div class="mb-4 bg-error text-error-content p-4 rounded-lg">
  <span class="text-xl">Неправильно!</span>
</div>
```

### **Progress & Status Components**

#### **Progress Indicator**
Floating progress badge with animated bar:
```svelte
<div class="absolute -top-10 left-1/2 -translate-x-1/2 z-10 bg-primary text-primary-content flex items-center gap-2 p-2 pl-4 pr-4 rounded-full translate-y-5 -rotate-2 text-main-font text-xl">
  <progress class="progress progress-neutral w-56" value={progress} max="100"></progress>
  {currentQuestionNumber}&nbsp;из&nbsp;{total}
</div>
```

#### **Loading States**
Consistent loading indicators:
```svelte
<!-- Full screen loading -->
<div class="absolute top-0 left-0 z-50 w-full h-full flex items-center backdrop-blur-sm">
  <div class="text-center mx-auto flex gap-2">
    <span class="loading loading-spinner loading-lg text-primary"></span>
    <p class="text-lg text-base-content/70">Загрузка...</p>
  </div>
</div>

<!-- Button loading state -->
<button class="btn btn-secondary loading">
  <span class="loading loading-spinner loading-md"></span>
  Загрузка...
</button>
```

### **Navigation Components**

#### **Back Button**
Telegram-aware navigation:
```svelte
<!-- Fallback for non-Telegram environments -->
<button class="btn btn-ghost btn-sm flex items-center gap-2">
  <svg class="size-4"><!-- back arrow --></svg>
  назад
</button>
```

#### **Countdown Timer**
Visual countdown for auto-progression:
```svelte
<span class="countdown text-primary font-bold">
  <span style="--value:{countdown};" aria-live="polite">{countdown}</span>
</span>
```

### **Specialized Components**

#### **Quiz Question Display**
Complete question presentation with image support:
```svelte
<div class="text-primary-content relative">
  {#if question.image}
    <div class="mb-3 place-self-center">
      <img src={question.image} alt="Quiz question illustration" 
           class="w-full mx-auto rounded-xl" loading="lazy" />
    </div>
  {/if}
  <h2 class="text-2xl font-semibold">{question.text}</h2>
</div>
```

#### **Result Display**
Game completion screen with score visualization:
```svelte
<div class="w-40 h-40 mask mask-squircle bg-gradient-to-b from-black/50 to-primary flex flex-col items-center justify-center">
  <div class="text-center">
    <div class="text-4xl font-bold text-main-font">{totalScore}</div>
    <div>
      <span class="text-xl font-bold text-main-font">из</span>
      <span class="text-2xl font-bold text-main-font">{total}</span>
    </div>
  </div>
</div>
```

## **Application Structure**

### **Page Layouts**

#### **Home Page (`/`)**
- **Container:** `max-w-md` centered layout
- **Background:** `bg-primary` full-screen
- **Sections:** 
  - App title with playful rotation effects
  - Daily quiz section with description and play button
  - Game ID input section with validation

#### **Game Page (`/game/[gameId]`)**
- **Container:** `max-w-md` with `p-2` padding
- **Components:** BackButton, QuizQuestion, QuizResult, Loading
- **States:** Loading, Question, Result, Error

#### **Layout (`+layout.svelte`)**
- **Background:** `bg-primary-content` full-screen
- **Features:** Telegram SDK integration, toast container
- **Responsive:** Mobile-first with `min-h-screen`

### **State Management**

#### **Quiz State Store**
Centralized state management for quiz progression:
- Current question data
- Progress tracking (answered/correct/total)
- Result data
- Loading and interaction states

#### **User Store**
Telegram user data management:
- User information from Telegram SDK
- Chat context and type
- Persistent storage

### **API Integration**

#### **Client Functions**
- `getGameState(gameId)` - Fetch current game state
- `submitAnswer(gameId, questionId, optionId)` - Submit quiz answer
- `resetGame(gameId)` - Reset game progress
- `getDailyGame()` - Get daily quiz game ID

### **Accessibility Features**

#### **Keyboard Navigation**
- Tab navigation for all interactive elements
- Enter/Space key support for buttons
- Escape key for dismissible elements

#### **Screen Reader Support**
- Proper ARIA labels and roles
- Live regions for dynamic content
- Semantic HTML structure

#### **Visual Accessibility**
- High contrast color combinations
- Focus indicators on interactive elements
- Loading states with appropriate feedback

### **Telegram Mini App Integration**

#### **SDK Features**
- Automatic initialization and user detection
- Native back button integration
- Share functionality with fallback
- Haptic feedback for interactions

#### **Platform Detection**
- Telegram-only access restriction
- Fallback screens for non-Telegram environments
- Deep linking support via start parameters

## **Development Guidelines**

### **Component Creation**
1. Use daisyUI components as base
2. Apply consistent `max-w-md` container pattern
3. Include proper TypeScript interfaces
4. Add accessibility attributes
5. Implement loading states

### **Styling Conventions**
1. Use utility-first approach with Tailwind
2. Apply `text-main-font` for branded typography
3. Use semantic color tokens (primary, secondary, etc.)
4. Maintain consistent spacing with `p-4`, `gap-4`
5. Add hover and transition effects

### **State Management**
1. Use Svelte stores for global state
2. Implement proper error handling
3. Add loading states for async operations
4. Maintain data consistency across components

### **Performance Optimization**
1. Lazy load images with `loading="lazy"`
2. Use `$derived` for computed values
3. Implement proper cleanup for timers/intervals
4. Optimize bundle size with tree shaking

## **File Structure Reference**

```
web/src/
├── lib/
│   ├── components/          # Reusable UI components
│   │   ├── QuizQuestion.svelte
│   │   ├── QuizResult.svelte
│   │   ├── BackButton.svelte
│   │   ├── Loading.svelte
│   │   └── ...
│   ├── stores/             # Global state management
│   │   ├── quiz.ts
│   │   └── user.ts
│   ├── api/                # API client functions
│   │   └── client.ts
│   ├── toast/              # Toast notification system
│   ├── actions/            # Svelte actions (feedback, etc.)
│   └── utils/              # Utility functions
├── routes/                 # SvelteKit routes
│   ├── +layout.svelte      # Global layout
│   ├── +page.svelte        # Home page
│   └── game/[gameId]/      # Dynamic game routes
├── app.css                 # Global styles and theme
└── app.html                # HTML template
```

This design system ensures consistency, accessibility, and maintainability across the entire quiz application while providing an optimal user experience within the Telegram Mini App environment.