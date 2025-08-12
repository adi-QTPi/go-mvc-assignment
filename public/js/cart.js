const to_menu_page = JSON.parse(sessionStorage.getItem("to_menu_page"));
const item_in_cart = JSON.parse(sessionStorage.getItem("item_in_cart"));

async function render_cart(item_in_cart){
    const cart_space = document.getElementsByClassName("cart_space")[0];
    cart_space.innerHTML = ``;
    for (let items of item_in_cart){
        let new_el = document.createElement("div");
        new_el.classList.add("card", "d-flex", "flex-row", "col-12", "col-lg-10", "mx-auto");
        let sub_total = items.quantity * items.price;

        new_el.innerHTML = `
            <div class="card-body flex-grow">
                <div class="card-title fs-2">
                    ${items.item_name}
                </div>
                <div class="card-subtitle text-muted form-floating col-12">
                    <input type="text" class="form-control item_instruction_input" data-item-id="${items.item_id}" placeholder="abc" id="instruction">
                    <label for="instruction">Instruction for the Chef</label>
                </div>
            </div>
            <div class="flex-shrink-1 d-flex flex-column me-2 align-items-center justify-content-center col-4">
                <div class="fs-3 sub_total">₹${sub_total}</div>
                <div class = "d-flex flex-row align-items-center gap-2">
                    <button type="button" class="btn btn-dark qty-btn minus-btn btn-sm" data-item-id="${items.item_id}">-</button>
                    <span class="qty-display fs-2">${items.quantity}</span>
                    <button type="button" class="btn btn-dark qty-btn plus-btn btn-sm" data-item-id="${items.item_id}">+</button>
                </div>
            </div>
            <div class="flex-shrink-1 d-flex align-items-center justify-content-center mx-1 me-3">
                <button type="button" class="delete-btn btn btn-danger">Remove</button>
            </div>
        `;
        cart_space.appendChild(new_el);

        const minus_button = new_el.querySelector(".minus-btn");
        const plus_button = new_el.querySelector(".plus-btn");
        const delete_button = new_el.querySelector(".delete-btn");
        const quantity_space = new_el.getElementsByClassName("qty-display")[0];
        const subtotal_space = new_el.getElementsByClassName("sub_total")[0];

        minus_button.addEventListener("click", () => {
            if (items.quantity > 1){
                items.quantity--;
                quantity_space.innerText = items.quantity; 
                subtotal_space.innerText = `₹${items.quantity * items.price}`;
            }
            else {
                quantity_space.innerText = items.quantity;
                subtotal_space.innerText = `₹${items.quantity * items.price}`;
            }
        });
        
        plus_button.addEventListener("click", () => {
            items.quantity++;
            quantity_space.innerText = items.quantity;
            subtotal_space.innerText = `₹${items.quantity * items.price}`;
        });

        delete_button.addEventListener("click", ()=>{
            const item_index = item_in_cart.findIndex(item => item.item_id === items.item_id);

            if(item_index !== -1){
                item_in_cart.splice(item_index,1);
                sessionStorage.setItem("item_in_cart", JSON.stringify(item_in_cart));
                new_el.remove();
            }
        })
    }
}

render_cart(item_in_cart);



const cart_form = document.getElementsByClassName("cart_form")[0];
cart_form.addEventListener("submit", (e)=>{
    e.preventDefault();

    const item_instruction_input = document.getElementsByClassName("item_instruction_input");

    for (let input of item_instruction_input){
        const item_id = Number(input.getAttribute("data-item-id"));
        const instruction_from_form = input.value.trim();

        const item_index = item_in_cart.findIndex(item => item.item_id === item_id);

        if(item_index !== -1){
            item_in_cart[item_index].instruction = instruction_from_form;
        }
    }

    sessionStorage.setItem("item_in_cart", JSON.stringify(item_in_cart));

    fetch("/api/order", {
    method: "POST",
    body: JSON.stringify(item_in_cart),
    headers: {
        "Content-type": "application/json; charset=UTF-8"
    }
    })
    .then((response) => response.json())
    .then((data)=>{
        window.location.href="/static/order"
    });
})