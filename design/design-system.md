# Design System — hello-word-22

> Source of truth: the approved `index.html` (preview: design preview URL from workspace).
> Every value below is extracted from it. Changing a value here without
> changing the approved design is a defect.

Last updated: 2025-08-28

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#FFFFFF` | Page background |
| `--color-text` | `#000000` | Body text |

#### Contrast audit

Every text-on-background pair actually used. Body text ≥ 4.5:1, large text (≥ 18.66px bold or ≥ 24px) ≥ 3:1, UI borders ≥ 3:1.

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21:1` | AA / AA Large |

### 1.2 Spacing

Base unit: `24px`. Every margin, padding, and gap in the product uses one of these.

| Token | Value |
|---|---|
| `--space-6` | `24px` |

### 1.3 Typography

Font families (include the fallback stack and how the font is loaded):

- Body: `Arial, Helvetica, sans-serif` (system fallback stack, no font load)
- Headings: `Arial, Helvetica, sans-serif` (same as body)
- Mono: not used

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-3xl` | `clamp(2rem, 5vw, 4rem)` | `1` | `400` | Greeting heading |

Heading levels are used in order and never skipped for visual sizing.

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--radius-sm` | not used | — |
| `--radius-md` | not used | — |
| `--radius-lg` | not used | — |
| `--radius-full` | not used | — |
| `--border-width` | not used | — |
| `--shadow-sm` | not used | — |
| `--shadow-md` | not used | — |
| `--shadow-lg` | not used | — |
| `--duration-fast` | not used | — |
| `--duration-base` | not used | — |
| `--easing` | not used | — |

Motion respects `prefers-reduced-motion: reduce`: state changes remain, movement is removed.

### 1.5 Layout and breakpoints

| Name | Min width | Container | Columns | Gutter |
|---|---|---|---|---|
| `sm` | not used | — | — | — |
| `md` | not used | — | — | — |
| `lg` | not used | — | — | — |
| `xl` | not used | — | — | — |

Z-index scale (only these values are allowed):

| Layer | Value |
|---|---|
| Base | `0` |
| Sticky header | not used |
| Dropdown | not used |
| Modal backdrop | not used |
| Modal | not used |
| Toast | not used |

## 2. Components

One subsection per reusable component. Every component lists **all** states.

### 2.1 Greeting screen

**Purpose** — Static centered message area for one-line greeting. Use only for this page; not for interactive content.

**Anatomy** — `[main] [h1]`

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-bg`, `--color-text`, `--space-6`, `--text-3xl` | One centered greeting page |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | viewport height | `24px` | `--text-3xl` |

**States** — every row must be filled in.

| State | Visual change | Tokens |
|---|---|---|
| Default | White page, black centered text | `--color-bg`, `--color-text` |
| Hover | None | None |
| Focus (keyboard) | None; no interactive target | None |
| Active / pressed | None | None |
| Disabled | None | None |
| Loading | None | None |
| Error | None | None |
| Empty | None; empty not used on this static view | None |

**Accessibility** — landmark `main` with `aria-label="Greeting screen"`; single `h1`; minimum hit target not applicable because no controls.

## 3. Content and formatting

- Voice and tone in one line: plain, neutral, no decoration.
- Date, time, number, and currency formats, with locale: not used.
- Capitalization rule for buttons, headings, and labels: heading keeps exact copy casing; no buttons or labels.
- Empty-state and error-message wording pattern: not used.

## 4. Known deviations

Places where the approved design does not follow its own rules or the
anti-patterns in `references/ai-defaults.md`. Record, do not silently fix.

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| `1.2 Spacing` | Only `24px` spacing token exists; no full scale | Approved page needs only one padding value | Add more tokens only if later screens need them |
| `1.4 Radius, border, shadow, motion` | No radius, border, shadow, or motion values used | Approved design is plain static text | Add only if future UI introduces surfaces or controls |
| `1.5 Layout and breakpoints` | No breakpoint/container system used | Single full-screen page does not need it | Add only if layout becomes responsive beyond default flow |
| `2.1 Greeting screen` | Component has no interactive states beyond default content | Product has no interactive controls | Add interaction states only if product gains controls |

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2025-08-28 | Initial design system for minimal greeting page | pending |
