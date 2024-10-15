document.getElementById('quoteRequest').addEventListener('submit', function(e) {
    e.preventDefault();
    const currency = document.getElementById('currency').value;
    const amount = document.getElementById('amount').value;

    fetch('/quote/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ currency, amount })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('quoteResponse').innerHTML = JSON.stringify(data);
    })
    .catch(error => {
        console.error('Error:', error);
    });
});
