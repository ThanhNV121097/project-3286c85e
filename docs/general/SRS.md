# SRS — General

Module: `general`
Last updated: 2025-02-14
Design: [View the approved design](http://localhost:8080/design/3286c85e-0326-405a-9b83-af1fdb928206)
Design system: `design/design-system.md`

## 1. Purpose

`general` module delivers one minimal end-to-end greeting page. It exists so the project can prove frontend, backend, database, and deployment wiring all work together. Without it, there is no visible proof that stored data can reach the page.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Guest | Not signed in visitor | View greeting page |

## 3. Scope

**In scope** — the functions specified below, by their plan titles:

- Show stored greeting page

**Out of scope** — name what a reader would reasonably expect here and say where it lives instead.

- Editing the greeting — not built in this project.
- Multiple pages, navigation, or authentication — not part of this slice.
- Design changes — belong to `design/index.html` and `design/design-system.md`, not this SRS.

## 4. Functional requirements

### 4.1 Show stored greeting page

**Requirement GENERAL-001 — Stored greeting reaches page**

*As a* guest, *I want to* open the page and see the greeting stored by the system, *so that* I can confirm the page is driven by backend data.

Behaviour:

1. When the guest opens the page, the system shows one centered line of text on a plain white background.
2. The text shown on the page comes from data stored by the system, not from frontend source.
3. The displayed text matches the stored greeting value exactly when the page loads.
4. The page remains visually plain: black text, white background, no animation, no extra content.

**Acceptance criteria** — each maps one-to-one onto a test case in `docs/general/test-cases/show-stored-greeting-page.md`.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | stored greeting value is `Hello Word` | guest opens page | page shows `Hello Word` |
| AC-2 | stored greeting value is changed to another text | guest opens page | page shows changed stored text |
| AC-3 | page is loaded | guest views screen | text is centered horizontally and vertically on white background with black text |

**Failure, boundary and permission behaviour**

| Case | Condition | Expected behaviour |
|---|---|---|
| Permission | Not applicable: page is public and has no sign-in state | Any guest can view same screen |
| Boundary | Stored greeting text is empty | Not applicable: approved design shows one greeting line only; empty state is not part of design |
| Failure | Backend or database unavailable | Not applicable in approved design; no error state is drawn on screen |
| Conflict | Stored greeting changes during view | Not applicable: this function is a read-only page load, and approved design shows one state only |

**Data touched** — the fields this function reads and writes, in product terms.

| Field | Type | Required | Rule |
|---|---|---|---|
| greeting text | text | yes | One stored greeting value, shown exactly as saved |

### 4.2 Dependencies and assumptions

- **Depends on:** backend API, for reading greeting text.
- **Depends on:** PostgreSQL, for storing one greeting row.
- **Assumption:** the stored greeting row exists before page load. If it does not, the project cannot show the greeting without adding a new empty-state decision.

| Open question | Proposed default | Who decides |
|---|---|---|
| Should the API return anything beyond the greeting text? | No, only the greeting text needed for this page | Stakeholder / TL |

## 5. Screens

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Greeting screen | Main centered greeting page | GENERAL-001 | default |

## 6. Non-functional requirements

| Area | Requirement |
|---|---|
| Performance | Page renders the greeting text within 1 second after response arrives |
| Accessibility | Greeting text contrast is at least 4.5:1 against background |
| Responsive | Screen stays centered at 320px width and above without horizontal scroll |

## 7. Dependencies and assumptions

- **Depends on:** backend API, for reading the greeting text.
- **Depends on:** PostgreSQL, for storing the single greeting row.
- **Assumption:** no loading, empty, or error state is part of the approved design.
- **Assumption:** white background `#FFFFFF` and black text `#000000` remain fixed for this module.

| Open question | Proposed default | Who decides |
|---|---|---|
| Should the greeting ever be editable through this project? | No | Stakeholder / TL |

## 8. Traceability

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Show stored greeting page | GENERAL-001 | `test-cases/show-stored-greeting-page.md` |
