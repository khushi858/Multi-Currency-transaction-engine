document.getElementById("payment-form").addEventListener("submit", async function(event) {
    event.preventDefault();

    const amount = document.getElementById("amount").value;
    const currency = document.getElementById("currency").value;

    const response = await fetch('/api/pay', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ amount, currency }),
    });

    const data = await response.json();
    const resultDiv = document.getElementById("result");

    if (data.success) {
        resultDiv.textContent = `Payment of ${data.convertedAmount} ${data.currency} was successful!`;
    } else {
        resultDiv.textContent = `Payment failed: ${data.message}`;
    }
});
