new Vue({
    el: '#app',
    data: {
        gameId: null,
        guess: null,
        message: '',
        messageClass: 'info',
        attempts: [],
        attempt: 0
    },
    created() {
        this.startGame();
    },
    methods: {
        startGame() {
            fetch('/start')
            .then(response => response.json())
            .then(data => {
                this.gameId = data.game_id;
                this.message = 'Введите число от 1 до 100';
                this.messageClass = 'info';
                this.attempt = 0;
                this.attempts = [];
                this.guess = null;
            })
            .catch(error => {
                console.error('Ошибка:', error);
                this.message = 'Ошибка при запуске игры';
                this.messageClass = 'error';
            });
        },
        submitGuess() {
            if (!this.guess || this.guess < 1 || this.guess > 100) {
                this.message = 'Пожалуйста, введите число от 1 до 100';
                this.messageClass = 'error';
                return;
            }

            fetch(`/guess?game_id=${this.gameId}&guess=${this.guess}`)
            .then(response => response.json())
            .then(data => {
                this.attempts.push({
                    guess: this.guess,
                    result: data.message
                });

                this.message = data.message;
                this.attempt = data.attempts;

                if (data.won) {
                    this.messageClass = 'success';
                } else if (data.remaining_attempts === 0) {
                    this.messageClass = 'error';
                } else {
                    this.messageClass = 'info';
                }

                this.guess = null;

                if (data.won || data.remaining_attempts === 0) {
                    setTimeout(() => {
                        this.startGame();
                    }, 3000);
                }
            })
            .catch(error => {
                console.error('Ошибка:', error);
                this.message = 'Ошибка при отправке предположения';
                this.messageClass = 'error';
            });
        }
    }
});
