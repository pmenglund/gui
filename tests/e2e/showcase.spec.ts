import { expect, test } from "@playwright/test";

test("overview links to every showcase section", async ({ page }) => {
  await page.goto("/");

  await expect(page.getByRole("heading", { name: "gomponents UI for HTMX-heavy Go apps" })).toBeVisible();
  await expect(page.getByRole("link", { name: "Browse primitives" })).toHaveAttribute("href", "/primitives");
  await expect(page.getByRole("link", { name: "Inspect forms" })).toHaveAttribute("href", "/forms");
  await expect(page.getByRole("link", { name: "Try interactions" })).toHaveAttribute("href", "/interactive");
  await expect(page.getByRole("link", { name: "Open HTMX flows" })).toHaveAttribute("href", "/htmx");
});

test("primitives page renders the full primitive catalog correctly", async ({ page }) => {
  await page.goto("/primitives");

  await expect(page.getByRole("heading", { name: "Primitive catalog" })).toBeVisible();
  await expect(page.getByRole("button", { name: "Default" })).toBeVisible();
  await expect(page.getByRole("button", { name: "Outline" })).toBeVisible();
  await expect(page.getByRole("button", { name: "Delete" })).toBeVisible();
  await expect(page.getByRole("button", { name: "Search" })).toBeVisible();
  await expect(page.getByText("Stable", { exact: true })).toBeVisible();
  await expect(page.getByText("Healthy", { exact: true })).toBeVisible();

  await expect(page.getByLabel("Project name")).toHaveAttribute("placeholder", "Aurora");
  await expect(page.getByLabel("Project summary")).toContainText("Server-rendered interfaces with a narrow client runtime.");

  await expect(page.getByLabel("Enable shipping updates")).toBeChecked();
  await expect(page.getByRole("radio").nth(1)).toBeChecked();
  await expect(page.getByLabel("Compact mode")).toBeChecked();
  await expect(page.getByRole("combobox")).toContainText("Choose a region");

  await expect(page.getByRole("alert")).toContainText("Heads up");
  await expect(page.getByRole("status")).toContainText("Loading metrics");
  await expect(page.getByLabel("Priya Mehta")).toContainText("PM");
});

test("forms page renders wiring, navigation, and empty state correctly", async ({ page }) => {
  await page.goto("/forms");

  const email = page.getByLabel("Work email");
  await expect(email).toHaveAttribute("aria-invalid", "true");
  await expect(email).toHaveAttribute("aria-describedby", /field-work-email-description/);
  await expect(page.getByText("Please use a company email address.")).toBeVisible();

  await expect(page.getByRole("navigation", { name: "Breadcrumb" })).toContainText("Deployments");
  await expect(page.getByRole("table")).toContainText("api");
  await expect(page.getByRole("navigation", { name: "Pagination" })).toContainText("3");

  await expect(page.getByText("Everything is quiet")).toBeVisible();
  await expect(page.getByRole("button", { name: "Create runbook" })).toBeVisible();
});

test("interactive widgets respond to clicks and keyboard controls", async ({ page }) => {
  await page.goto("/interactive");

  await page.getByRole("button", { name: "Open dialog" }).click();
  await expect(page.getByText("The dialog content stays in place in the DOM and relies on focus restoration instead of portals.")).toBeVisible();
  await page.keyboard.press("Escape");
  await expect(page.getByText("The dialog content stays in place in the DOM and relies on focus restoration instead of portals.")).toBeHidden();

  await page.getByRole("button", { name: "Quick actions" }).click();
  await expect(page.getByText("View logs")).toBeVisible();
  await page.mouse.click(5, 5);
  await expect(page.getByText("View logs")).toBeHidden();

  await page.getByRole("tab", { name: "Overview" }).focus();
  await page.keyboard.press("ArrowRight");
  await expect(page.getByRole("tab", { name: "Activity" })).toBeFocused();
  await expect(page.getByText("The tabs runtime handles keyboard navigation with arrow keys.")).toBeVisible();

  await page.getByRole("button", { name: "Show toast" }).click();
  await expect(page.getByText("Settings saved")).toBeVisible();

  await page.getByRole("button", { name: "Open sheet" }).click();
  await expect(page.getByText("This panel uses the right edge of the viewport and the same close affordances.")).toBeVisible();
  await page.keyboard.press("Escape");
  await expect(page.getByText("This panel uses the right edge of the viewport and the same close affordances.")).toBeHidden();
});

test("htmx flows update fragments and swapped overlays stay interactive", async ({ page }) => {
  await page.goto("/htmx");

  await page.getByRole("button", { name: "Refresh counter fragment" }).click();
  await expect(page.getByText("Counter update 2")).toBeVisible();

  const email = page.getByLabel("Email");
  await email.click();
  await email.type("ops@example.com");
  await expect(page.locator("#email-hint")).toHaveText("Looks good for a company email.");

  await page.getByRole("link", { name: "2" }).last().click();
  await expect(page.getByText("Search")).toBeVisible();

  await page.getByRole("button", { name: "Load interactive fragment" }).click();
  await page.getByRole("button", { name: "Open swapped dialog" }).click();
  await expect(page.getByText("Delegated runtime hooks mean the client code does not need a manual rebind.")).toBeVisible();
});
