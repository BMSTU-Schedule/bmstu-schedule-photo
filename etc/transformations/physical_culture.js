function () {
    var spans = document.getElementsByTagName('span')
    for (var i = 0; i < spans.length; i++) {
        if (spans[i].textContent === 'Элективный курс по физической культуре и спорту' ||
            spans[i].textContent === 'Физ воспитание' ||
            spans[i].textContent === 'Физическая культура и спорт'
        ) {
            row = spans[i].parentNode.parentNode
            time_cell = row.querySelector('.bg-grey.text-nowrap')
            time_cell.innerText = {
                '08': '08:15 - 09:45',
                '10': '10:00 - 11:30',
                '12': '12:20 - 13:50',
                '13': '14:05 - 15:35',
                '15': '15:50 - 17:20',
                '17': '17:25 - 18:55',
            }[time_cell.innerText.substring(0, 2)]
        }
    }
}