(function () {
  var FOCUSABLE_SELECTOR =
    "[autofocus], button:not([disabled]), [href], input:not([disabled]), select:not([disabled]), textarea:not([disabled]), [tabindex]:not([tabindex='-1'])";

  function closestController(el) {
    return el ? el.closest("[data-ui-controller]") : null;
  }

  function contentFor(root) {
    return root ? root.querySelector("[data-ui-content]") : null;
  }

  function triggerFor(root) {
    return root ? root.querySelector("[data-ui-trigger]") : null;
  }

  function isFocusable(el) {
    return !!el && typeof el.matches === "function" && el.matches(FOCUSABLE_SELECTOR);
  }

  function firstFocusableWithin(el) {
    return el ? el.querySelector(FOCUSABLE_SELECTOR) : null;
  }

  function focusableTriggerFor(root) {
    var trigger = triggerFor(root);
    if (!trigger) return null;
    return isFocusable(trigger) ? trigger : firstFocusableWithin(trigger);
  }

  function openFocusTarget(root) {
    var content = contentFor(root);
    if (!content) return null;

    if (root.dataset.uiController === "dialog" || root.dataset.uiController === "sheet") {
      return content.querySelector("[role='dialog']");
    }

    if (root.dataset.uiController === "dropdownmenu") {
      return firstFocusableWithin(content);
    }

    return null;
  }

  function focusElement(el) {
    if (el && typeof el.focus === "function") {
      el.focus();
    }
  }

  function focusLater(el) {
    if (!el || typeof el.focus !== "function") return;

    if (typeof window.requestAnimationFrame === "function") {
      window.requestAnimationFrame(function () {
        el.focus();
      });
      return;
    }

    window.setTimeout(function () {
      el.focus();
    }, 0);
  }

  function restoreFocus(root) {
    focusLater(focusableTriggerFor(root));
  }

  function close(root) {
    setState(root, false);
    restoreFocus(root);
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
      focusElement(openFocusTarget(root));
    } else {
      restoreFocus(root);
    }
  }

  function syncTabs(root) {
    if (!root || root.dataset.uiController !== "tabs") return;
    var buttons = root.querySelectorAll("[data-ui-trigger][data-ui-target]");
    var panels = root.querySelectorAll("[data-ui-content][data-ui-target]");
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

    var closeButton = event.target.closest("[data-ui-close]");
    if (closeButton) {
      var closeRoot = closestController(closeButton);
      close(closeRoot);
      return;
    }

    document.querySelectorAll("[data-ui-controller='dropdownmenu'][data-ui-state='open']").forEach(function (root) {
      if (!root.contains(event.target)) {
        close(root);
      }
    });
  });

  document.addEventListener("keydown", function (event) {
    if (event.key === "Escape") {
      document.querySelectorAll("[data-ui-controller][data-ui-state='open']").forEach(function (root) {
        close(root);
      });
    }

    var root = closestController(event.target);
    if (!root || root.dataset.uiController !== "tabs") return;

    var buttons = Array.from(root.querySelectorAll("[data-ui-trigger][data-ui-target]"));
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
