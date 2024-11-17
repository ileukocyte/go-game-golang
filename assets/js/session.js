const N = 9; // TODO: fetch from db when a session is created

document.addEventListener('DOMContentLoaded', () => {
    const boardGridBody = document.querySelector('#board-grid tbody');
    const stoneGridBody = document.querySelector('#stone-grid tbody');

    let black = true;

    document.getElementById('stone-grid').style.transform = `scale(${(N + 1)/N})`;

    // Generate NxN grid
    for (let i = 0; i < N; i++) {
        const row = document.createElement('tr');

        for (let j = 0; j < N; j++) {
            const cell = document.createElement('td');
            row.appendChild(cell);
        }

        boardGridBody.appendChild(row);
    }

    // Generate (N+1)x(N+1) grid
    for (let i = 0; i < N + 1; i++) {
        const row = document.createElement('tr');

        for (let j = 0; j < N + 1; j++) {
            const cell = document.createElement('td');
            row.appendChild(cell);
        }

        stoneGridBody.appendChild(row);
    }

    stoneGridBody.addEventListener('click', (event) => {
        if (event.target.tagName === 'TD' && !event.target.classList.contains('stone')) {
            let cell = event.target;

            let row = cell.parentElement.rowIndex;
            let col = cell.cellIndex;

            cell.classList.add('stone', black ? 'black-stone' : 'white-stone');
            console.log(`Placed stone at (${row}, ${col})`);

            // TODO: handle the move on the backend

            black = !black;
        }
    });
});