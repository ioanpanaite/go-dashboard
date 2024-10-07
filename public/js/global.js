let count = 0;

// Always escape HTML for text arguments!
function escapeHtml(html) {
  const div = document.createElement("div");
  div.textContent = html;
  return div.innerHTML;
}
// Custom function to emit toast notifications
function notify(
  message,
  variant = "primary",
  icon = "info-circle",
  duration = 3000,
) {
  const alert = Object.assign(document.createElement("sl-alert"), {
    variant,
    closable: true,
    duration: duration,
    innerHTML: `
            <sl-icon name="${icon}" slot="icon"></sl-icon>
            ${escapeHtml(message)}
          `,
  });

  document.body.append(alert);
  return alert.toast();
}

function toggleDetails(rowId) {
  var detailsRow = document.getElementById(rowId);
  detailsRow.classList.toggle("hidden");
}

function toggleChargesVisibility(employeeId) {
  var chargesRows = document.querySelectorAll(
    ".chargedefEmployee-" + employeeId,
  );
  var toggleIcon = document.getElementById("toggleIcon-" + employeeId);

  var hide = chargesRows[0].classList.contains("hidden");
  chargesRows.forEach((row) => {
    row.classList.toggle("hidden");
  });

  // Update the icon based on visibility
  if (hide) {
    toggleIcon.classList.remove("fa-chevron-right");
    toggleIcon.classList.add("fa-chevron-down");
  } else {
    toggleIcon.classList.remove("fa-chevron-down");
    toggleIcon.classList.add("fa-chevron-right");
  }
}

function hideVisibility(iconId, targetQuery) {
  var elements = document.querySelectorAll(targetQuery);
  var toggleIcon = document.getElementById("toggleIcon-" + iconId);
  toggleIcon.classList.remove("fa-chevron-down");
  toggleIcon.classList.add("fa-chevron-right");

  elements.forEach((element) => {
    // Initialize a counter attribute if not already set
    if (!element.hasAttribute("data-hidden-counter")) {
      element.setAttribute("data-hidden-counter", 0);
    }
    element.classList.add("hidden");
  });
}


function toggleVisibility(iconId, targetQuery) {
  var elements = document.querySelectorAll(targetQuery);
  var toggleIcon = document.getElementById("toggleIcon-" + iconId);
  var hide = toggleIcon.classList.contains("fa-chevron-down");

  // Update the icon based on visibility
  if (hide) {
    toggleIcon.classList.remove("fa-chevron-down");
    toggleIcon.classList.add("fa-chevron-right");
  } else {
    toggleIcon.classList.remove("fa-chevron-right");
    toggleIcon.classList.add("fa-chevron-down");
  }

  //console.log("elements");
  //console.log(elements);
  elements.forEach((element) => {
    // Initialize a counter attribute if not already set
    if (!element.hasAttribute("data-hidden-counter")) {
      element.setAttribute("data-hidden-counter", 0);
    }
    var counter = parseInt(element.getAttribute("data-hidden-counter"));
    // Check if the item is already hidden
    var isRowHidden = element.classList.contains("hidden");
    // If showing, check if the counter is not 0 before removing "hidden" class
    if (hide) {
      if (isRowHidden) {
        counter++;
      } else {
        counter = 0;
        element.classList.add("hidden");
      }
    } else {
      if (counter === 0) {
        element.classList.remove("hidden");
      }
      if (counter > 0) {
        counter--;
      }
    }
    // Update the counter attribute
    element.setAttribute("data-hidden-counter", counter);
  });
}

////////////////////////////////
// Manage chip selectors code
//
//
////////////////////////////////
function initializeDropdown(containerId) {
  const container = document.getElementById(containerId);

  // Select elements based on their roles or tags
  const dropdown = container.querySelector("sl-dropdown");
  const inputField = container.querySelector("sl-input");
  const menu = container.querySelector("sl-menu");
  const selectionDiv = container.querySelector(".chipsholder");
  const hiddenInput = container.querySelector('input[type="hidden"]');

  // Attach event listeners
  inputField.addEventListener("sl-input", () => {
    dropdown.show();
  });

  menu.addEventListener("sl-select", (event) => {
    const selectedItem = event.detail.item;
    const newTag = createTag(selectedItem);
    selectionDiv.appendChild(newTag);
    updateHiddenInput(hiddenInput, selectionDiv);
  });

  document.addEventListener("sl-remove", (event) => {
    // Get all sl-tag elements inside selectionDiv
    const tags = selectionDiv.querySelectorAll("sl-tag");

    for (let tag of tags) {
      // Check if the current tag is the event target or an ancestor of the event target
      if (tag === event.target || tag.contains(event.target)) {
        tag.remove();
        updateHiddenInput(hiddenInput, selectionDiv);
        break; // Once the matching tag is found and removed, exit the loop
      }
    }
  });
}

function createTag(selectedItem) {
  const newTag = document.createElement("sl-tag");
  newTag.setAttribute("removable", "");
  newTag.setAttribute("name", selectedItem.getAttribute("value"));
  newTag.setAttribute("value", selectedItem.getAttribute("value"));
  newTag.setAttribute("data-uuid", selectedItem.getAttribute("data-uuid"));
  newTag.setAttribute("data-id", selectedItem.getAttribute("data-id"));
  newTag.textContent = selectedItem.textContent.trim();
  return newTag;
}

function updateHiddenInput(hiddenInput, selectionDiv) {
  const tags = selectionDiv.querySelectorAll("sl-tag");
  const values = Array.from(tags).map((tag) => tag.getAttribute("value"));
  hiddenInput.value = values.join(",");
}

// Function to show the dropdown menu
function showDropdownMenu(menuId) {
  const menu = document.getElementById(menuId);
  menu.show();
}

function extractInputsToJson() {
  const controlsDiv = document.getElementById("controls");
  const inputs = controlsDiv.querySelectorAll("input");

  const state = Array.from(inputs).reduce((acc, input) => {
    acc[input.id || input.name] =
      input.type === "checkbox" ? input.checked : input.value;
    return acc;
  }, {});

  filters = state;

  return state;
}

function isNonEmptyState(state) {
  return Object.values(state).some((value) => value !== "" && value !== false);
}

function saveState() {
  const state = extractInputsToJson();
  // Save the state of chips
  const chipElements = document.querySelectorAll(
    "#searchCompany-FILTER sl-tag",
  );
  if (chipElements.length > 0) {
    state.companyIds = Array.from(chipElements).map((tag) =>
      tag.getAttribute("value"),
    );
    state.companyNames = Array.from(chipElements).map((tag) => tag.textContent);
    state.companyIds = state.companyIds.join(",");
    state.companyNames = state.companyNames.join(",");
  }

  // Save the state of chips
  const chipEmployees = document.querySelectorAll(
    "#searchEmployee-FILTER sl-tag",
  );
  if (chipEmployees.length > 0) {
    state.employeeIds = Array.from(chipEmployees).map((tag) =>
      tag.getAttribute("value"),
    );
    state.employeeNames = Array.from(chipEmployees).map(
      (tag) => tag.textContent,
    );
    state.employeeIds = state.employeeIds.join(",");
    state.employeeNames = state.employeeNames.join(",");
  }

  filterState = state;

  localStorage.setItem("inputStates", JSON.stringify(state));
  // Save the state of the filters
  updateURLWithFilterState();
  updateAllTables();
}

function loadState() {
  const queryParams = new URLSearchParams(window.location.search);
  let state = {};

  // Check if there are any URL parameters
  if (queryParams.toString() !== "") {
    for (const [key, value] of queryParams) {
      state[key] = value === "true" ? true : value === "false" ? false : value;
    }
  } else {
    // Fallback to local storage if no URL parameters
    const savedState = localStorage.getItem("inputStates");
    if (savedState) {
      state = JSON.parse(savedState);
    }
  }

  // Apply the state to inputs
  if (isNonEmptyState(state)) {
    //console.log(state);

    const controlsDiv = document.getElementById("controls");
    const inputs = controlsDiv.querySelectorAll("input");

    filters = state; // Assuming 'filters' is a global variable you're using elsewhere

    inputs.forEach((input) => {
      const key = input.id || input.name;
      if (key && state.hasOwnProperty(key)) {
        if (input.type === "checkbox") {
          input.checked = state[key];
        } else {
          input.value = state[key];
        }
      }
    });

    // Load company chips
    if (state.companyIds && state.companyNames) {
      const chipValuesArray = state.companyIds.split(",");
      const chipNamesArray = state.companyNames.split(",");
      chipValuesArray.forEach((value, index) => {
        const newTag = document.createElement("sl-tag");
        newTag.setAttribute("removable", "");
        newTag.setAttribute("value", value);
        newTag.textContent = chipNamesArray[index];
        document
          .getElementById("searchCompany-FILTER")
          .querySelector(".chipsholder")
          .appendChild(newTag);
      });
    }

    // Load employee chips
    if (state.employeeIds && state.employeeNames) {
      const chipValuesArrayE = state.employeeIds.split(",");
      const chipNamesArrayE = state.employeeNames.split(",");
      chipValuesArrayE.forEach((value, index) => {
        const newTag = document.createElement("sl-tag");
        newTag.setAttribute("removable", "");
        newTag.setAttribute("value", value);
        newTag.textContent = chipNamesArrayE[index];
        document
          .getElementById("searchEmployee-FILTER")
          .querySelector(".chipsholder")
          .appendChild(newTag);
      });
    }

    // // Check if 'seePending' is part of the loaded state and update the checkbox
    // if (state.hasOwnProperty("seePending")) {
    //   //console.log("seePending is part of the loaded state");
    //   //console.log(state.seePending);
    //   //const checkbox = document.getElementById("seePending");
    //   //checkbox.checked = state.seePending;
    //   //const faCheckmark = document.getElementById("faCheckmark");
    //   //const chip = document.getElementById("chip");

    //   // Update the visual state based on the checkbox state
    //   if (state.seePending) {
    //     faCheckmark.classList.remove("hidden");
    //     chip.classList.add("bg-green-300");
    //     chip.classList.remove("bg-gray-200");
    //   } else {
    //     faCheckmark.classList.add("hidden");
    //     chip.classList.remove("bg-green-300");
    //     chip.classList.add("bg-gray-200");
    //   }
    // }

    updateAllTables(); // Update tables or other components as needed
    saveState(); // Save the state loaded from URL to local storage
    triggerFilterUpdate(); // Call the function to update filters
  }
}

function clearInputsAndStorage() {
  const controlsDiv = document.getElementById("controls");
  const inputs = controlsDiv.querySelectorAll("input");

  // Resetting each input field
  inputs.forEach((input) => {
    if (input.type === "checkbox") {
      input.checked = false;
    } else {
      input.value = "";
    }
  });

  // Select all chip containers
  const chipsContainers = document.querySelectorAll(".chipsholder");
  console.log("chipsContainers");
  console.log(chipsContainers);
  // Remove all sl-tag elements within each chip container
  chipsContainers.forEach((container) => {
    console.log(container.firstChild);
    while (container.firstChild) {
      container.removeChild(container.firstChild);
    }
  });

  // Clearing local storage
  localStorage.removeItem("inputStates");

  updateAllTables(); // Update tables or other components as needed
  saveState(); // Save the state loaded from URL to local storage
}

function updateHtmxHeaders(element, newData) {
  // Parse existing headers
  let existingHeaders = {};
  try {
    existingHeaders = JSON.parse(element.getAttribute("hx-headers"));
  } catch (e) {
    //console.error("Error parsing existing HTMX headers:", e);
  }

  // Update with new data
  let updatedHeaders = { ...existingHeaders, ...newData };

  // Set updated headers back on the element
  element.setAttribute("hx-headers", JSON.stringify(updatedHeaders));
}

function updateHtmxParams(element, newData) {
  // Parse existing headers
  let existingHeaders = {};
  try {
    existingHeaders = JSON.parse(element.getAttribute("hx-headers"));
  } catch (e) {
    //console.error("Error parsing existing HTMX headers:", e);
  }

  // Update with new data
  let updatedHeaders = { ...existingHeaders, ...newData };

  // Convert array values to comma-separated strings
  for (let key in updatedHeaders) {
    if (Array.isArray(updatedHeaders[key])) {
      updatedHeaders[key] = updatedHeaders[key].join(",");
    }
  }

  // Set updated headers back on the element
  element.setAttribute("hx-vals", JSON.stringify(updatedHeaders));
}

function updateHtmxUrlsWithQueryParams(element, newData) {
  // Base URL extraction (remove existing query parameters if any)
  let baseUrl = element.getAttribute("hx-get").split("?")[0];

  // Merge newData with existing query parameters
  let existingParams = new URLSearchParams(
    element.getAttribute("hx-get").split("?")[1] || "",
  );
  Object.keys(newData).forEach((key) => {
    existingParams.set(key, newData[key]);
  });

  // Construct the new URL with updated query parameters
  let newUrl = `${baseUrl}?${existingParams.toString()}`;

  // Set the new URL in the hx-get attribute
  element.setAttribute("hx-get", newUrl);

  //console.log("Updated HTMX URL:", newUrl);
}

function updateURLWithFilterState() {
  const queryParams = new URLSearchParams();

  for (const key in filterState) {
    // Only add parameter if it's not empty
    if (filterState[key] !== "" && !(filterState[key] === false)) {
      // This checks for non-empty strings and true values for checkboxes
      queryParams.set(key, filterState[key]);
    }
  }

  // Update URL without reloading the page, only if queryParams is not empty
  if (queryParams.toString() !== "") {
    history.pushState({}, "", "?" + queryParams.toString());
  } else {
    // Clear the query parameters from the URL if there are no filters
    history.pushState({}, "", window.location.pathname);
  }
}

function updateAllTables() {
  // Update global filter data
  // updateGlobalFilterData();
  // Get all tables that need to be updated
  var tables = document.querySelectorAll(".filterable-hx-rows"); // Assuming a common class for tables
  // Update each table
  tables.forEach(function (table) {
    updateHtmxParams(table, filters);
  });
  //console.log(filters);
  //console.log("Updated all tables");
}

function triggerFilterUpdate() {
  //console.log("Triggering filter update");
  tbodies = document.querySelectorAll(".filterable-hx-rows");

  //console.log(tbodies);
  tbodies.forEach(function (element) {
    element.dispatchEvent(new Event("filterUpdateEvent"));
  });
}


//// CUSTOM FUNCTIONS
// function getInputValues(containerId) {
//   var inputs = document.querySelectorAll('#' + containerId + ' input');
//   var inputValues = {};

//   inputs.forEach(function(input) {
//       inputValues[input.name] = input.value;
//   });

//   return inputValues;
// }

// function addHeaders(){
// var newData = getInputValues('filtersBox');
// elements = document.querySelectorAll(".add-headers");

// elements.forEach(function(element) {
//     // Parse existing headers
// let existingHeaders = {};
// try {
//   existingHeaders = JSON.parse(element.getAttribute("hx-headers"));
// } catch (e) {
//   existingHeaders ={}
//   //console.error("Error parsing existing HTMX headers:", e);
// }

// // Update with new data
// let updatedHeaders = { ...existingHeaders, ...newData };

// // Set updated headers back on the element
// element.setAttribute("hx-headers", JSON.stringify(updatedHeaders));
// });
// }
function cancelRequest(requestId) {
  var request = htmx.find(requestId);
  if (request && request.xhr) {
    request.xhr.abort();
  }
}

function formatDate(date) {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0'); // Add leading zero if month is single digit
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

function getUUIDs(selData) {
  // get uuid from each selected row
  var dataUuids = []
  selData.forEach(el => {
    dataUuids.push(el.getAttribute('data-uuid'));
  });
  return dataUuids;
}



function getSelectedCharges() {
  var companyTable = document.querySelectorAll('#companyCharges tbody tr.selected');
  var employeeTable = document.querySelectorAll('#employeeCharges tbody tr.selected');
  var payrollTable = document.querySelectorAll('#pemployeeCharges tbody tr.selected');

  var employeeTableUUIDs = new Set();
  document.getElementById("employeeTableIds").dataset.index.split(",").forEach(function (tableId) {
    var table = document.querySelectorAll(`#${tableId} tbody tr.selected`);
    var pTable = document.querySelectorAll(`#${tableId} tbody tr.selected`);
    var pUuid = getUUIDs(pTable)
    var uuids = getUUIDs(table);

    uuids.forEach(uuid => employeeTableUUIDs.add(uuid));
    pUuid.forEach(uuid => employeeTableUUIDs.add(uuid))
  });

  var cUuids = getUUIDs(companyTable);
  var eUuids = getUUIDs(employeeTable);
  var pUuids = getUUIDs(payrollTable);

  // Convert Set back to an array to maintain the order of elements
  var uniqueUUIDs = Array.from(employeeTableUUIDs);

  // Concatenate all UUID arrays and remove duplicates
  var result = [...new Set([...cUuids, ...eUuids, ...pUuids, ...uniqueUUIDs])];

  return result;
}
