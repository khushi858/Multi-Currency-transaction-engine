document.addEventListener("DOMContentLoaded", async () => {
    const currencySelects = document.querySelectorAll("select");
    const convertButton = document.getElementById("convertButton");
    const resultDiv = document.getElementById("result");

    // Fetch currency data
    async function fetchCurrencies() {
        const response = await fetch("http://localhost:8080/api/currencies");
        const currencies = await response.json();
        currencies.forEach(currency => {
            const optionFrom = document.createElement("option");
            optionFrom.value = currency.code;
            optionFrom.textContent = currency.code;

            const optionTo = optionFrom.cloneNode(true);

            currencySelects[0].appendChild(optionFrom);
            currencySelects[1].appendChild(optionTo);
        });
    }

    // Convert currency
    async function convertCurrency() {
        const amount = document.getElementById("amount").value;
        const fromCurrency = document.getElementById("fromCurrency").value;
        const toCurrency = document.getElementById("toCurrency").value;

        const response = await fetch(`http://localhost:8080/api/convert?amount=${amount}&from=${fromCurrency}&to=${toCurrency}`);
        const result = await response.json();
        resultDiv.textContent = `${amount} ${fromCurrency} = ${result.convertedAmount} ${toCurrency}`;
    }

    convertButton.addEventListener("click", convertCurrency);
    await fetchCurrencies();
});
