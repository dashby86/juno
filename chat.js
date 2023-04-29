function createInputForm() {
    // Create form elements
    const form = document.createElement('form');
    const label = document.createElement('label');
    const input = document.createElement('input');
    const button = document.createElement('button');

    // Set attributes and content
    label.setAttribute('for', 'inputData');
    label.textContent = 'Enter data:';
    input.type = 'text';
    input.id = 'inputData';
    input.name = 'inputData';
    button.type = 'submit';
    button.textContent = 'Submit';

    // Add elements to the form
    form.appendChild(label);
    form.appendChild(input);
    form.appendChild(button);

    // Add form to the DOM
    document.body.appendChild(form);

    // Attach event listener to the form
    form.addEventListener('submit', async (event) => {
        event.preventDefault();

        const data = {
            value: input.value
        };

        try {
            const response = await fetch('https://your-api-endpoint.example.com', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            });

            if (response.ok) {
                console.log('Data posted successfully:', data);
            } else {
                console.error('Error posting data:', response);
            }
        } catch (error) {
            console.error('Error posting data:', error);
        }
    });
}

// Call the function when the DOM content is loaded
document.addEventListener('DOMContentLoaded', createInputForm);
