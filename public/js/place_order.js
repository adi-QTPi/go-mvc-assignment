let item_in_cart = [];

const metaTag = document.querySelector('meta[name="menu-data"]');
const to_menu_page = JSON.parse(metaTag.getAttribute('content'));

console.log(to_menu_page)

const go_to_cart_space = document.getElementsByClassName("go-to-cart-space")[0];
const go_to_cart_space_text = document.getElementsByClassName("go-to-cart-space-text")[0];

document.addEventListener("click", function(event) {
    if (event.target.classList.contains("add-to-cart")) {
        let id = Number(event.target.id);
        if (item_in_cart.find((item) => item.item_id === id)) {
            const item_index = item_in_cart.findIndex(item => item.item_id === id);
            event.target.innerText = "add to cart";
            event.target.classList.remove("btn-dark");
            event.target.classList.add("btn-danger");
            item_in_cart.splice(item_index, 1);
        } else {
            const foundItem = to_menu_page.ItemSlice.find(el => el.item_id === id);
            if (foundItem) {
                let new_item_for_cart = {
                    item_id: foundItem.item_id,
                    item_name: foundItem.item_name,
                    price: parseInt(foundItem.price, 10),
                    quantity: 1
                }
                event.target.innerText = "Remove";
                event.target.classList.remove("btn-danger");
                event.target.classList.add("btn-dark")
                item_in_cart.push(new_item_for_cart);
            }
        }
        toggle_add_to_cart_button_label();
        toggle_to_cart_button_visibility();
        update_text_in_element(go_to_cart_space_text, `You have <span class="caveat-cursive fs-1"> ${item_in_cart.length} </span> item(s) in your cart !`);
    }
});

toggle_to_cart_space_visibility("none");

let selected_filters = []; let filtered_menu = [];
const filter_buttons = document.getElementsByClassName("filter-buttons")[0].children;

for (let filter of filter_buttons){
    filter.addEventListener("click", ()=>{
        const cat_id = Number(filter.getAttribute("data-btn"));
        if(filter.classList.contains("btn-filter")){
            filter.classList.remove("btn-filter");
            filter.classList.add("btn-filter-clicked");
            selected_filters.push(cat_id);
        }
        else{
            filter.classList.add("btn-filter");
            filter.classList.remove("btn-filter-clicked");
            selected_filters.splice(selected_filters.indexOf(cat_id), 1);
        }
        create_filtered_menu(selected_filters);
        render_filtered_menu(filtered_menu);
        toggle_add_to_cart_button_label();
    })
}

const clear_filter_button = document.getElementsByClassName("clear-filter-button")[0];
clear_filter_button.addEventListener("click", ()=>{
    for(let el of filter_buttons){
        if(!el.classList.contains("btn-filter")){
            el.classList.remove("btn-filter-clicked");
            el.classList.add("btn-filter");
        }
        selected_filters = [];
    }
    create_filtered_menu(selected_filters);
    render_filtered_menu(filtered_menu);
    toggle_add_to_cart_button_label();
})

async function toggle_to_cart_button_visibility(){
    if(!item_in_cart.length){

        toggle_to_cart_space_visibility("none");
    }
    else{

        toggle_to_cart_space_visibility("block");
    }
}

async function toggle_add_to_cart_button_label(){
    let add_to_cart_buttons = document.getElementsByClassName("add-to-cart");
    for(let buttons of add_to_cart_buttons){
        let id = Number(buttons.id);
        buttons.classList.remove("btn-danger", "btn-dark");
        buttons.innerText = "";
        if(item_in_cart.some((item)=> item.item_id === id)){
            buttons.classList.add("btn-dark");
            buttons.innerText = "Remove from Cart";
        }
        else {
            buttons.classList.add("btn-danger");
            buttons.innerText = "Add to Cart"
        }
    }
}

async function toggle_to_cart_space_visibility(vis_value){
    for(let ch of go_to_cart_space.children){
        ch.style.display = vis_value;
    }
}
async function update_text_in_element(element, html){
    element.innerHTML = html;
}

async function create_filtered_menu(selected_filters){
    filtered_menu = [];

    for(let cat_id of selected_filters){
        const target_item_array = to_menu_page.ItemSlice.filter(item => item.cat_id === cat_id || item.subcat_id === cat_id);
        for(let target_item of target_item_array){
            if(target_item && !filtered_menu.some(item => item.item_id === target_item.item_id)){
                filtered_menu.push(target_item);
            }
        }
    }
}

const filtered_menu_space = document.getElementsByClassName("filtered-menu-space")[0];
filtered_menu_space.style.display = "none";
async function render_filtered_menu(filtered_menu){
    if(filtered_menu_space.style.display === "none"){
        filtered_menu_space.style.display = "block";
    }
    if(!filtered_menu.length){
        filtered_menu_space.style.display = "none";
    }
    filtered_menu_space.innerHTML = "";

    let menu_space = document.createElement("div");
    menu_space.classList.add("d-flex", "flex-column", "flex-lg-row", "flex-wrap", "align-items-center", "justify-content-center", "m-1", "m-md-4", "gap-2", "gap-md-4")
    
    let reverse_filtered_menu = [...filtered_menu].reverse();

    for( let items of reverse_filtered_menu){
        let new_el = document.createElement("div");
        new_el.classList.add("card" ,"col-10", "col-lg-5" ,"d-flex" ,"flex-column");

        let img_path = "/public/images/sample_food.png";
        if(items.display_pic.Valid ){
            img_path = items.display_pic.String;
        }

        if(to_menu_page.XUser.role === "admin"){
            new_el.innerHTML = `
            <div class="ratio ratio-21x9 menu-card-image-container">
                <img class="card-img-top menu-card-image" src=
                "${img_path}"
                alt="sample-pic">
            </div>
            <div class="d-flex flex-row">
                <div class="card-body flex-grow">
                    <div class="card-title fs-2">
                        ${items.item_name}
                    </div>
                    <div class="card-subtitle text-muted">
                        wait time : <span class="text-queen-pink">
                            ${items.cook_time_min}
                        </span> Minutes
                    </div>
                    <div class="d-flex flex-row col-10">
                        <div class="flex-fill border p-2 m-1 text-center truculenta-normal fs-5">
                            ${items.category}
                        </div>
                        <div class="flex-fill border p-2 m-1 text-center truculenta-normal fs-5">
                            ${items.subcategory}
                        </div>
                    </div>
                </div>
                <div class="flex-shrink-1 d-flex flex-column me-2 align-items-center justify-content-center">
                    <div class="fs-3">₹ ${items.price}</div>
                    <form action="/api/item/d/${items.item_id}"  method="post">
                        <button type="submit" class="btn btn-danger">Delete Item</button>
                    </form>
                </div>
            </div>
            `;
        }
        else if (to_menu_page.XUser.role === "cook") {
            new_el.innerHTML = `
            <div class="ratio ratio-21x9 menu-card-image-container">
                <img class="card-img-top menu-card-image" src=
                "${img_path}"
                alt="sample-pic">
            </div>
            <div class="d-flex flex-row">
                <div class="card-body flex-grow">
                    <div class="card-title fs-2">
                        ${items.item_name}
                    </div>
                    <div class="card-subtitle text-muted">
                        wait time : <span class="text-queen-pink">
                            ${items.cook_time_min}
                        </span> Minutes
                    </div>
                    <div class="d-flex flex-row col-10">
                        <div class="flex-fill border p-2 m-1 text-center truculenta-normal fs-5">
                            ${items.category}
                        </div>
                        <div class="flex-fill border p-2 m-1 text-center truculenta-normal fs-5">
                            ${items.subcategory}
                        </div>
                    </div>
                </div>
                <div class="flex-shrink-1 d-flex flex-column me-2 align-items-center justify-content-center col-2">
                    <div class="fs-3 mx-auto">₹ ${items.price}</div>
                </div>
            </div>
            `;
        }
        else{
            new_el.innerHTML = `
            <div class="ratio ratio-21x9 menu-card-image-container">
                <img class="card-img-top menu-card-image" src=
                "${img_path}"
                alt="sample-pic">
            </div>
            <div class="d-flex flex-row">
                <div class="card-body flex-grow">
                    <div class="card-title fs-2">
                        ${items.item_name}
                    </div>
                    <div class="card-subtitle text-muted">
                        wait time : <span class="text-queen-pink">
                            ${items.cook_time_min}
                        </span> Minutes
                    </div>
                    <div class="d-flex flex-row col-10">
                        <div class="flex-fill border p-2 m-1 text-center truculenta-normal fs-5">
                            ${items.category}
                        </div>
                        <div class="flex-fill border p-2 m-1 text-center truculenta-normal fs-5">
                            ${items.subcategory}
                        </div>
                    </div>
                </div>
                <div class="flex-shrink-1 d-flex flex-column me-2 align-items-center justify-content-center">
                    <div class="fs-3">₹ ${items.price}</div>
                    <button class="add-to-cart btn btn-danger" id="${items.item_id}">Add to Cart</button>
                </div>
            </div>
            `;
        }
        menu_space.appendChild(new_el);
    }
    filtered_menu_space.appendChild(menu_space);
}

const to_cart_button = document.getElementsByClassName("to_cart_button")[0];
to_cart_button.addEventListener("click", ()=>{
    sessionStorage.setItem("to_menu_page", JSON.stringify(to_menu_page));
    sessionStorage.setItem("item_in_cart", JSON.stringify(item_in_cart));
    window.location.href="/static/cart";
})
