// Function to open the popup
function openPopup() {
  document.getElementById("popup").style.display = "block";
}

// Function to close the popup
function closePopup() {
  document.getElementById("popup").style.display = "none";
}

// Fetch data from the API and render products
async function fetchAndRenderProducts() {
  const productsContainer = document.getElementById("products-container");

  try {
    const response = await fetch("http://localhost:9999/get");
    const data = await response.json();

    data.forEach((product) => {
      const productElement = document.createElement("div");
      productElement.className = "product";

      const productImage = document.createElement("img");
      productImage.src = product.imgurl;
      console.log(product.imgurl);
      productElement.appendChild(productImage);

      const productName = document.createElement("h2");
      productName.textContent = product.name;
      productElement.appendChild(productName);

      const productDescription = document.createElement("p");
      productDescription.textContent = product.description;
      productElement.appendChild(productDescription);

      const productPrice = document.createElement("p");
      productPrice.textContent = "Price: $" + product.price;
      productElement.appendChild(productPrice);

      const buyNowButton = document.createElement("button");
      buyNowButton.className = "buy-now-btn";
      buyNowButton.textContent = "Buy Now";
      buyNowButton.addEventListener("click", () =>
        openOrderPopup(product.name)
      );
      productElement.appendChild(buyNowButton);

      productsContainer.appendChild(productElement);
    });
  } catch (error) {
    console.log("Error:", error);
  }
}

// Open the order form popup with the selected item
function openOrderPopup(selectedItem) {
  const selectedItemInput = document.getElementById("selected-item");
  selectedItemInput.value = selectedItem;
  openPopup();
}

// Get the device's location
function getLocation() {
  if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(showPosition);
  } else {
    alert("Geolocation is not supported by your browser.");
  }
}

// Show the device's position
function showPosition(position) {
  const latitude = position.coords.latitude;
  const longitude = position.coords.longitude;
  reverseGeocode(latitude, longitude);
}

// Reverse geocode the coordinates to get the location
function reverseGeocode(latitude, longitude) {
  const apiKey = "ef5e77a3d1304c0398673dc981410731"; // Replace with your OpenCage API key
  const url = `https://api.opencagedata.com/geocode/v1/json?key=${apiKey}&q=${latitude}+${longitude}`;

  fetch(url)
    .then(response => response.json())
    .then(data => {
      if (data.results && data.results.length > 0) {
        const location = data.results[0].formatted;
        console.log("Location:", location);
        // Update the address field with the location name
        document.getElementById("address").value = location;
      } else {
        console.log("Location not found.");
      }
    })
    .catch(error => {
      console.log("Error fetching location:", error);
    });
}

// Submit the order form
function submitOrderForm(event) {
  event.preventDefault();

  // Retrieve form values
  const selectedItem = document.getElementById("selected-item").value;
  const quantity = document.getElementById("quantity").value;
  const name = document.getElementById("name").value;
  const phone = document.getElementById("phone").value;
  const address = document.getElementById("address").value;

  // Do something with the form data (e.g., send it to the server)
  console.log("Selected Item:", selectedItem);
  console.log("Quantity:", quantity);
  console.log("Name:", name);
  console.log("Phone Number:", phone);
  console.log("Address:", address);

  // Close the popup
  closePopup();
}

// Add submit event listener to the order form
document.getElementById("order-form").addEventListener("submit", submitOrderForm);

// Fetch and render products when the page loads
window.onload = fetchAndRenderProducts;
