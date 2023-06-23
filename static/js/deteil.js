
    // Assuming you have an API endpoint that returns the data in JSON format
    const apiUrl = 'http://localhost:9999/getall';

    // Fetch data from the API
    fetch(apiUrl)
      .then(response => response.json())
      .then(data => {
        const tableBody = document.getElementById('table-body');

        // Iterate over each item and create a table row
        data.forEach(item => {
          const row = document.createElement('tr');
          const { id, imgurl, name, price, description } = item;

          // Create table cells for each item property
          const idCell = document.createElement('td');
          idCell.textContent = id;

          const imgCell = document.createElement('td');
          const img = document.createElement('img');
          img.src = imgurl;
          imgCell.appendChild(img);

          const nameCell = document.createElement('td');
          nameCell.textContent = name;

          const priceCell = document.createElement('td');
          priceCell.textContent = price;

          const descriptionCell = document.createElement('td');
          descriptionCell.textContent = description;

          const actionsCell = document.createElement('td');
          const updateBtn = document.createElement('button');
          updateBtn.classList.add('update-btn');
          updateBtn.textContent = 'Update';

          const deleteBtn = document.createElement('button');
          deleteBtn.classList.add('delete-btn');
          deleteBtn.textContent = 'Delete';

          // Add click event listeners to update and delete buttons
          updateBtn.addEventListener('click', () => {
            openUpdatePopup(item); // Pass the item object to the update popup
          });

          deleteBtn.addEventListener('click', () => {
            const itemId = item.id; // Get the ID of the selected item

            // Prompt the user for confirmation before deleting the item
            const confirmDelete = confirm('Are you sure you want to delete this item?');
            if (confirmDelete) {
              // Make a DELETE request to the API endpoint with the item ID
              const deleteUrl = `http://localhost:9999/items/${itemId}`;
              fetch(deleteUrl, {
                method: 'DELETE',
              })
                .then(response => {
                  // Check if the request was successful
                  if (response.ok) {
                    // Remove the item row from the table
                    tableBody.removeChild(row);
                  } else {
                    // Handle error response
                    console.log('Delete request failed');
                  }
                })
                .catch(error => {
                  console.log('Error:', error);
                });
            }
          });

          actionsCell.appendChild(updateBtn);
          actionsCell.appendChild(deleteBtn);

          // Append all cells to the row
          row.appendChild(idCell);
          row.appendChild(imgCell);
          row.appendChild(nameCell);
          row.appendChild(priceCell);
          row.appendChild(descriptionCell);
          row.appendChild(actionsCell);

          // Append the row to the table body
          tableBody.appendChild(row);
        });
      })
      .catch(error => {
        console.log('Error:', error);
      });

    // Function to open the update popup with pre-filled values
    function openUpdatePopup(item) {
      const popup = document.getElementById('update-popup');
      const updateImageInput = document.getElementById('update-image');
      const updateDescriptionInput = document.getElementById('update-description');
      const updateNameInput = document.getElementById('update-name');
      const updatePriceInput = document.getElementById('update-price');
      const updateIdInput = document.getElementById('update-id');
      const updateSubmitBtn = document.getElementById('update-submit');

      // Set the pre-filled values in the input fields
      updateImageInput.value = item.imgurl;
      updateDescriptionInput.value = item.description;
      updateNameInput.value = item.name;
      updatePriceInput.value = item.price;
      updateIdInput.value = item.id;

      // Add click event listener to the update submit button
      updateSubmitBtn.addEventListener('click', () => {
        const updatedImage = updateImageInput.value;
        const updatedDescription = updateDescriptionInput.value;
        const updatedName = updateNameInput.value;
        const updatedPrice = updatePriceInput.value;
        const id = item.id;

        // Handle the update functionality with the updated values and id
        // Add your logic here

        // Close the update popup after updating the value
        popup.style.display = 'none';
      });

      // Display the update popup
      popup.style.display = 'block';
    }