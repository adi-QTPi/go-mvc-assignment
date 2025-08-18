function save_cart() {
    localStorage.setItem("item_in_cart", JSON.stringify(item_in_cart));
    localStorage.setItem("to_menu_page", JSON.stringify(to_menu_page));
}

function load_cart() {
    const item_in_storage = localStorage.getItem("item_in_cart");
    if (item_in_storage) {
        return JSON.parse(item_in_storage);
    } else {
        return [];
    }
}

item_in_cart = load_cart();

const metaTag = document.querySelector('meta[name="menu-data"]');
const to_menu_page = JSON.parse(metaTag.getAttribute("content"));
const go_to_cart_space = document.getElementsByClassName("go-to-cart-space")[0];
const go_to_cart_space_text = document.getElementsByClassName(
    "go-to-cart-space-text"
)[0];

document.addEventListener("click", function (event) {
    let parent = event.target.parentElement;
    if (parent.classList.contains("add-to-cart")) {
        let new_target = parent.querySelector(".add-to-cart-btn");
        let id = Number(new_target.id);
        if (item_in_cart.find((item) => item.item_id === id)) {
            const item_index = item_in_cart.findIndex((item) => item.item_id === id);
            item_in_cart.splice(item_index, 1);
        } else {
            const foundItem = to_menu_page.ItemSlice.find((el) => el.item_id === id);
            if (foundItem) {
                let new_item_for_cart = {
                    item_id: foundItem.item_id,
                    item_name: foundItem.item_name,
                    price: parseInt(foundItem.price, 10),
                    quantity: 1,
                };
                item_in_cart.push(new_item_for_cart);
            }
        }
        save_cart();
        toggle_add_to_cart_button_label();
        toggle_to_cart_button_visibility();
        update_text_in_element(
            go_to_cart_space_text,
            `You have <span class="caveat-cursive fs-1"> ${item_in_cart.length} </span> item(s) in your cart !`
        );
    }
});

toggle_to_cart_button_visibility();

update_text_in_element(
    go_to_cart_space_text,
    `You have <span class="caveat-cursive fs-1"> ${item_in_cart.length} </span> item(s) in your cart !`
);

let selected_filters = [];
let filtered_menu = [];
const filter_buttons =
    document.getElementsByClassName("filter-buttons")[0].children;

for (let filter of filter_buttons) {
    filter.addEventListener("click", () => {
        const cat_id = Number(filter.getAttribute("data-btn"));
        if (filter.classList.contains("btn-filter")) {
            filter.classList.remove("btn-filter");
            filter.classList.add("btn-filter-clicked");
            selected_filters.push(cat_id);
        } else {
            filter.classList.add("btn-filter");
            filter.classList.remove("btn-filter-clicked");
            selected_filters.splice(selected_filters.indexOf(cat_id), 1);
        }
        create_filtered_menu(selected_filters);
        renderWithSearchResults();
        toggle_add_to_cart_button_label();
    });
}

const clear_filter_button = document.getElementsByClassName(
    "clear-filter-button"
)[0];
clear_filter_button.addEventListener("click", () => {
    for (let el of filter_buttons) {
        if (!el.classList.contains("btn-filter")) {
            el.classList.remove("btn-filter-clicked");
            el.classList.add("btn-filter");
        }
        selected_filters = [];
    }
    search_input.value = "";
    create_filtered_menu(selected_filters);
    renderWithSearchResults();
    toggle_add_to_cart_button_label();
});

async function toggle_add_to_cart_button_label() {
    let add_to_cart_buttons = document.getElementsByClassName("add-to-cart");
    item_in_cart = load_cart();
    for (let button of add_to_cart_buttons) {
        let id = Number(button.id);
        let buttons = button.querySelector(".add-to-cart-btn");
        if (item_in_cart.some((item) => item.item_id == id)) {
            const foundItem = item_in_cart.find((item) => item.item_id == id);
            let quantity = foundItem.quantity;

            button.innerHTML = `<div class="d-flex flex-column align-items-center justify-content-center">
                <div class="">
                    Qty : <span id="item${id}" class="qty-display fs-3 truculenta-normal">${quantity}</span>
                </div>
                <div class="btn-group " role="group" aria-label="Quantity controls">
                    <button type="button" id="${id}" class="qty-minus btn bg-light-pink rounded-left-3 fs-4">-</button>
                    <button type="button" id="${id}" class="qty-remove btn bg-turq-bold text-white fs-4">x</button>
                    <button type="button" id="${id}" class="qty-plus btn bg-light-pink rounded-right-3 fs-4">+</button>
                </div>
            </div>`;
        }
        else {
            button.innerHTML = `<button
                  class="add-to-cart-btn btn bg-queen-pink text-white rounded-5 p-2 m-1 fs-4"
                  id="${id}"
                >
                  Add to Cart
                </button>`;
        }
    }
}

document.addEventListener("click", function (event) {
    const clickedElement = event.target;
    if (clickedElement.classList.contains('qty-minus')) {
        const itemId = clickedElement.id;
        decrementQty(itemId);
    } else if (clickedElement.classList.contains('qty-plus')) {
        const itemId = clickedElement.id;
        incrementQty(itemId);
    } else if (clickedElement.classList.contains('qty-remove')) {
        const itemId = clickedElement.id;
        removeItem(itemId);
    }
});

function incrementQty(itemId) {
    const itemToIncrement = item_in_cart.find((item) => item.item_id == itemId);
    itemToIncrement.quantity++;
    save_cart();
    update_qty_display(itemToIncrement.item_id, itemToIncrement.quantity);
}
function decrementQty(itemId) {
    const itemToDecrement = item_in_cart.find((item) => item.item_id == itemId);
    if (itemToDecrement.quantity > 1) {
        itemToDecrement.quantity--;
    } else {
        removeItem(itemId);
        save_cart();
        return;
    }
    save_cart();
    update_qty_display(itemToDecrement.item_id, itemToDecrement.quantity);
}
function removeItem(itemId) {
    const itemIndex = item_in_cart.findIndex((item) => item.item_id == itemId);
    if (itemIndex > -1) {
        item_in_cart.splice(itemIndex, 1);
    }
    save_cart();
    toggle_add_to_cart_button_label();
    toggle_to_cart_button_visibility();
    update_text_in_element(
        go_to_cart_space_text,
        `You have <span class="caveat-cursive fs-1"> ${item_in_cart.length} </span> item(s) in your cart !`
    );
}
function update_qty_display(id, number) {
    const qty_display_space = document.querySelector(`#item${String(id)}.qty-display`);
    if (!qty_display_space) return;
    qty_display_space.innerText = number
}

async function toggle_to_cart_space_visibility(vis_value) {
    for (let ch of go_to_cart_space.children) {
        ch.style.display = vis_value;
    }
}
async function update_text_in_element(element, html) {
    element.innerHTML = html;
}

async function create_filtered_menu(selected_filters) {
    filtered_menu = [];

    for (let cat_id of selected_filters) {
        const target_item_array = to_menu_page.ItemSlice.filter(
            (item) => item.cat_id === cat_id || item.subcat_id === cat_id
        );

        for (let target_item of target_item_array) {
            if (
                target_item &&
                !filtered_menu.some((item) => item.item_id === target_item.item_id)
            ) {
                filtered_menu.push(target_item);
            }
        }
    }
}

const filtered_menu_space = document.getElementsByClassName(
    "filtered-menu-space"
)[0];
const full_menu_space = document.getElementsByClassName("full-menu-space")[0];

create_filtered_menu(selected_filters);
render_filtered_menu(filtered_menu);
toggle_add_to_cart_button_label();
toggle_to_cart_button_visibility();

filtered_menu_space.style.display = "none";

function generateMenuItemHtml(item, role) {
    let action_button_html = "";
    let img_path;
    if (item.display_pic.Valid) {
        img_path = item.display_pic.String;
    } else {
        img_path = "/public/images/sample_food.png";
    }

    if (role === "admin") {
        action_button_html = `
            <form action="/api/item/d/${item.item_id}" method="post">
                <button type="submit" class="btn bg-queen-pink text-white rounded-5 p-2 m-1 fs-4">Delete Item</button>
            </form>
        `;
    } else if (role === "cook") {
        action_button_html = `
            <div class="mx-auto"></div>
        `;
    } else {
        action_button_html = `
            <div class="qty-controls add-to-cart" id="${item.item_id}">
                <button
                  class="add-to-cart-btn btn bg-queen-pink text-white rounded-5 p-2 m-1 fs-4"
                  id="${item.item_id}"
                >
                  Add to Cart
                </button>
              </div>
        `;
    }

    return `
        <div class="ratio ratio-21x9 menu-card-image-container rounded-5">
                <img class="card-img-top menu-card-image rounded-top-5" src="${img_path}" alt="sample-pic">
            </div>
            <div class="d-flex flex-row">
                <div class="card-body flex-grow">
                    <div class="card-title fs-2">${item.item_name}</div>
                    <div class="card-subtitle text-muted">
                        wait time : <span class="text-queen-pink">${item.cook_time_min}</span> Minutes
                    </div>
                    <div class="d-flex flex-row col-10">
                        <div class="flex-fill border p-2 m-1 text-center truculenta-normal fs-5 rounded-5 bg-light-pink">${item.category}</div>
                        <div class="flex-fill border p-2 m-1 text-center truculenta-normal fs-5 rounded-5 bg-light-pink">${item.subcategory}</div>
                    </div>
                </div>
                <div class="d-flex flex-column me-2 align-items-center justify-content-around col-3">
                    <div class="fs-3">₹ ${item.price}</div>
                    ${action_button_html}
                </div>
            </div>
    `;
}

async function render_filtered_menu(filtered_menu) {
    if (filtered_menu_space.style.display === "none") {
        filtered_menu_space.style.display = "block";
        full_menu_space.style.display = "none";
    }
    if (!filtered_menu.length) {
        filtered_menu_space.style.display = "none";
        full_menu_space.style.display = "block";
    }
    filtered_menu_space.innerHTML = "";

    let menu_space = document.createElement("div");
    menu_space.classList.add(
        "d-flex",
        "flex-column",
        "flex-lg-row",
        "flex-wrap",
        "align-items-center",
        "justify-content-center",
        "m-1",
        "m-md-4",
        "gap-2",
        "gap-md-4"
    );

    let reverse_filtered_menu = [...filtered_menu].reverse();

    for (let items of reverse_filtered_menu) {
        let new_el = document.createElement("div");
        new_el.classList.add(
            "card",
            "col-10",
            "col-lg-5",
            "d-flex",
            "flex-column",
            "rounded-5"
        );

        let img_path = "/public/images/sample_food.png";
        if (items.display_pic.Valid) {
            img_path = items.display_pic.String;
        }

        new_el.innerHTML = generateMenuItemHtml(items, to_menu_page.XUser.role);

        menu_space.appendChild(new_el);
    }
    filtered_menu_space.appendChild(menu_space);
}

const search_input = document.getElementById("search-input");

search_input.addEventListener("input", (e) => {
    renderWithSearchResults();
});

async function renderWithSearchResults() {
    const searchTerm = search_input.value.trim();
    if (searchTerm === "") {
        create_filtered_menu(selected_filters);
        render_filtered_menu(filtered_menu);
    } else {
        const searchResults = to_menu_page.ItemSlice.filter((item) =>
            item.item_name.toLowerCase().includes(searchTerm.toLowerCase())
        );
        if (selected_filters.length > 0) {
            const filteredBySearchAndCategory = searchResults.filter(
                (item) =>
                    selected_filters.includes(item.cat_id) ||
                    selected_filters.includes(item.subcat_id)
            );

            if (filteredBySearchAndCategory.length === 0) {
                filtered_menu_space.innerHTML = "";
                let noItemsCard = document.createElement("div");
                noItemCardRender(noItemsCard);
                filtered_menu_space.appendChild(noItemsCard);
                filtered_menu_space.style.display = "block";
            } else {
                render_filtered_menu(filteredBySearchAndCategory);
            }
        } else {
            if (searchResults.length == 0) {
                filtered_menu_space.innerHTML = "";
                let noItemsCard = document.createElement("div");
                noItemCardRender(noItemsCard);

                filtered_menu_space.appendChild(noItemsCard);
                filtered_menu_space.style.display = "block";
            } else {
                render_filtered_menu(searchResults);
            }
        }
    }
    toggle_add_to_cart_button_label();
}

function noItemCardRender(element) {
    element.classList.add(
        "card",
        "col-12",
        "col-lg-8",
        "rounded-pill",
        "d-flex",
        "flex-column",
        "m-4",
        "mx-auto",
        "p-4",
        "text-center"
    );
    element.innerHTML = `
        <div class="card-body col-12"> 
            <div class="card-title fs-1 truculenta-normal">No Items Found</div>
            <div class="card-text fs-4 truculenta-normal">Try searching for something else or clearing your filters.</div>
        </div>
    `;
}

async function toggle_to_cart_button_visibility() {
    if (!item_in_cart.length) {
        toggle_to_cart_space_visibility("none");
    } else {
        toggle_to_cart_space_visibility("block");
    }
}
