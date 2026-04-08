(function () {
  function closestController(el) {
    return el ? el.closest("[data-ui-controller]") : null;
  }

  function ownedBy(root, el) {
    return closestController(el) === root;
  }

  function queryOwned(root, selector) {
    if (!root) return [];
    return Array.from(root.querySelectorAll(selector)).filter(function (el) {
      return ownedBy(root, el);
    });
  }

  var focusableSelector = "[autofocus], button, [href], input, select, textarea, [tabindex]:not([tabindex='-1'])";

  function contentFor(root) {
    return queryOwned(root, "[data-ui-content]")[0] || null;
  }

  function triggerFor(root) {
    var trigger = queryOwned(root, "[data-ui-trigger]")[0] || null;
    if (!trigger) return null;
    if (trigger.matches(focusableSelector)) return trigger;
    return trigger.querySelector(focusableSelector) || trigger;
  }

  function setState(root, open) {
    if (!root) return;
    root.dataset.uiState = open ? "open" : "closed";
    root.setAttribute("data-ui-state", root.dataset.uiState);

    var content = contentFor(root);
    if (content) {
      if (open) {
        content.removeAttribute("hidden");
        content.classList.remove("hidden");
      } else {
        content.setAttribute("hidden", "hidden");
        content.classList.add("hidden");
      }
    }

    root.dispatchEvent(
      new CustomEvent(open ? "ui:open" : "ui:close", {
        bubbles: true,
        detail: { controller: root.dataset.uiController },
      }),
    );
  }

  function toggle(root) {
    var open = root && root.dataset.uiState === "open";
    setState(root, !open);
    if (!open) {
      var focusTarget = contentFor(root);
      if (focusTarget) {
        var autoFocus = focusTarget.querySelector(focusableSelector);
        if (autoFocus) autoFocus.focus();
      }
    } else {
      var trigger = triggerFor(root);
      if (trigger) trigger.focus();
    }
  }

  function syncTabs(root) {
    if (!root || root.dataset.uiController !== "tabs") return;
    var buttons = queryOwned(root, "[data-ui-trigger][data-ui-target]");
    var panels = queryOwned(root, "[data-ui-content][data-ui-target]");
    var active = root.dataset.uiValue || (buttons[0] && buttons[0].dataset.uiTarget) || "";
    buttons.forEach(function (button) {
      var current = button.dataset.uiTarget === active;
      button.setAttribute("aria-selected", current ? "true" : "false");
      button.tabIndex = current ? 0 : -1;
    });
    panels.forEach(function (panel) {
      if (panel.dataset.uiTarget === active) {
        panel.removeAttribute("hidden");
        panel.classList.remove("hidden");
      } else {
        panel.setAttribute("hidden", "hidden");
        panel.classList.add("hidden");
      }
    });
  }

  function syncAll(scope) {
    (scope || document).querySelectorAll("[data-ui-controller]").forEach(function (root) {
      if (root.dataset.uiController === "tabs") {
        syncTabs(root);
        return;
      }
      setState(root, root.dataset.uiState === "open");
    });
  }

  document.addEventListener("click", function (event) {
    var trigger = event.target.closest("[data-ui-trigger]");
    if (trigger) {
      var root = closestController(trigger);
      if (root) {
        event.preventDefault();
        if (root.dataset.uiController === "tabs") {
          root.dataset.uiValue = trigger.dataset.uiTarget || "";
          syncTabs(root);
          root.dispatchEvent(new CustomEvent("ui:toggle", { bubbles: true, detail: { controller: "tabs", value: root.dataset.uiValue } }));
        } else {
          toggle(root);
        }
        return;
      }
    }

    var close = event.target.closest("[data-ui-close]");
    if (close) {
      var closeRoot = closestController(close);
      setState(closeRoot, false);
      var closeTrigger = triggerFor(closeRoot);
      if (closeTrigger) closeTrigger.focus();
      return;
    }

    document.querySelectorAll("[data-ui-controller='dropdownmenu'][data-ui-state='open']").forEach(function (root) {
      if (!root.contains(event.target)) {
        setState(root, false);
      }
    });
  });

  document.addEventListener("keydown", function (event) {
    if (event.key === "Escape") {
      document.querySelectorAll("[data-ui-controller][data-ui-state='open']").forEach(function (root) {
        setState(root, false);
        var trigger = triggerFor(root);
        if (trigger) trigger.focus();
      });
    }

    var root = closestController(event.target);
    if (!root || root.dataset.uiController !== "tabs") return;

    var buttons = queryOwned(root, "[data-ui-trigger][data-ui-target]");
    var index = buttons.indexOf(event.target);
    if (index === -1) return;

    var next = index;
    if (event.key === "ArrowRight") next = (index + 1) % buttons.length;
    if (event.key === "ArrowLeft") next = (index - 1 + buttons.length) % buttons.length;
    if (next !== index) {
      event.preventDefault();
      root.dataset.uiValue = buttons[next].dataset.uiTarget || "";
      syncTabs(root);
      buttons[next].focus();
    }
  });

  document.addEventListener("DOMContentLoaded", function () {
    syncAll(document);
  });

  document.body.addEventListener("htmx:afterSwap", function (event) {
    syncAll(event.target);
  });
})();
