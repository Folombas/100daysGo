// Modern JavaScript for Go Educational Module
class GoEducationalApp {
    constructor() {
        this.currentSection = 'overview';
        this.wsConnection = null;
        this.init();
    }

    init() {
        this.setupNavigation();
        this.loadUseCases();
        this.setupDemos();
        this.connectWebSocket();
        this.setupAnimations();
    }

    setupNavigation() {
        const navButtons = document.querySelectorAll('.nav-btn');
        const sections = document.querySelectorAll('.section');

        navButtons.forEach(btn => {
            btn.addEventListener('click', () => {
                const targetSection = btn.dataset.section;
                
                // Update active nav button
                navButtons.forEach(b => b.classList.remove('active'));
                btn.classList.add('active');
                
                // Update active section
                sections.forEach(s => s.classList.remove('active'));
                document.getElementById(targetSection).classList.add('active');
                
                this.currentSection = targetSection;
                
                // Load section-specific content
                this.loadSectionContent(targetSection);
            });
        });
    }

    async loadUseCases() {
        try {
            const response = await fetch('/api/usecases');
            const useCases = await response.json();
            this.renderUseCases(useCases);
        } catch (error) {
            console.error('Failed to load use cases:', error);
        }
    }

    renderUseCases(useCases) {
        const container = document.getElementById('usecases-container');
        container.innerHTML = '';

        useCases.forEach(useCase => {
            const card = document.createElement('div');
            card.className = 'usecase-card';
            card.innerHTML = `
                <div class="usecase-header">
                    <div class="usecase-category">${useCase.category}</div>
                    <h3 class="usecase-title">${useCase.title}</h3>
                    <p class="usecase-description">${useCase.description}</p>
                </div>
                <div class="usecase-example">
                    <h4>Example: ${useCase.example}</h4>
                    <div class="usecase-code">
                        <pre><code class="language-go">${useCase.code}</code></pre>
                    </div>
                </div>
            `;
            container.appendChild(card);
        });

        // Re-highlight code
        if (window.Prism) {
            Prism.highlightAll();
        }
    }

    setupDemos() {
        const demoButtons = document.querySelectorAll('.demo-btn');
        
        demoButtons.forEach(btn => {
            btn.addEventListener('click', async () => {
                const demoType = btn.dataset.demo;
                await this.runDemo(demoType);
            });
        });
    }

    async runDemo(demoType) {
        const resultDiv = document.getElementById(`${demoType}-result`);
        const btn = document.querySelector(`[data-demo="${demoType}"]`);
        
        // Show loading state
        btn.innerHTML = '<span class="loading"></span> Running...';
        btn.disabled = true;
        
        try {
            const response = await fetch(`/api/demo/${demoType}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                }
            });
            
            const result = await response.json();
            
            // Display result
            resultDiv.innerHTML = `
                <h4>${result.operation}</h4>
                <div class="demo-data">
                    ${Object.entries(result.data).map(([key, value]) => 
                        `<div class="data-item">
                            <span class="data-key">${key.replace(/_/g, ' ')}:</span>
                            <span class="data-value">${value}</span>
                        </div>`
                    ).join('')}
                </div>
                <div class="demo-timestamp">${new Date(result.timestamp).toLocaleTimeString()}</div>
            `;
            
            resultDiv.classList.add('show');
            
            // Add success animation
            resultDiv.style.animation = 'slideDown 0.3s ease-out';
            
        } catch (error) {
            resultDiv.innerHTML = `
                <h4>Error</h4>
                <p class="error">Failed to run demo: ${error.message}</p>
            `;
            resultDiv.classList.add('show');
        } finally {
            // Reset button
            btn.innerHTML = 'Run Demo';
            btn.disabled = false;
        }
    }

    connectWebSocket() {
        const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
        const wsUrl = `${protocol}//${window.location.host}/ws`;
        
        try {
            this.wsConnection = new WebSocket(wsUrl);
            
            this.wsConnection.onopen = () => {
                console.log('WebSocket connected');
            };
            
            this.wsConnection.onmessage = (event) => {
                const data = JSON.parse(event.data);
                this.updateRealTimeStats(data.data);
            };
            
            this.wsConnection.onclose = () => {
                console.log('WebSocket disconnected');
                // Reconnect after 3 seconds
                setTimeout(() => this.connectWebSocket(), 3000);
            };
            
            this.wsConnection.onerror = (error) => {
                console.error('WebSocket error:', error);
            };
            
        } catch (error) {
            console.error('Failed to connect WebSocket:', error);
        }
    }

    updateRealTimeStats(data) {
        const goroutinesEl = document.getElementById('goroutines');
        const memoryEl = document.getElementById('memory-usage');
        const cpuEl = document.getElementById('cpu-usage');
        
        if (goroutinesEl && data.active_goroutines) {
            goroutinesEl.textContent = data.active_goroutines;
            this.animateValue(goroutinesEl, data.active_goroutines);
        }
        
        if (memoryEl && data.memory_usage) {
            memoryEl.textContent = data.memory_usage;
        }
        
        if (cpuEl && data.cpu_usage) {
            cpuEl.textContent = data.cpu_usage;
        }
    }

    animateValue(element, newValue) {
        const currentValue = parseInt(element.textContent) || 0;
        const increment = (newValue - currentValue) / 20;
        let current = currentValue;
        
        const timer = setInterval(() => {
            current += increment;
            if ((increment > 0 && current >= newValue) || (increment < 0 && current <= newValue)) {
                current = newValue;
                clearInterval(timer);
            }
            element.textContent = Math.round(current);
        }, 50);
    }

    loadSectionContent(section) {
        switch (section) {
            case 'examples':
                this.loadCodeExamples();
                break;
            case 'demos':
                this.updateRealTimeStats({
                    active_goroutines: Math.floor(Math.random() * 100) + 50,
                    memory_usage: '15.2MB',
                    cpu_usage: '12.5%'
                });
                break;
        }
    }

    async loadCodeExamples() {
        try {
            const response = await fetch('/api/usecases');
            const useCases = await response.json();
            this.renderCodeExamples(useCases);
        } catch (error) {
            console.error('Failed to load code examples:', error);
        }
    }

    renderCodeExamples(useCases) {
        const container = document.getElementById('examples-container');
        container.innerHTML = '';

        useCases.forEach((useCase, index) => {
            const card = document.createElement('div');
            card.className = 'example-card';
            card.innerHTML = `
                <div class="example-header">
                    <h3 class="example-title">${useCase.title}</h3>
                    <p class="example-description">${useCase.description}</p>
                </div>
                <div class="example-code">
                    <pre><code class="language-go">${useCase.code}</code></pre>
                </div>
            `;
            container.appendChild(card);
        });

        // Re-highlight code
        if (window.Prism) {
            Prism.highlightAll();
        }
    }

    setupAnimations() {
        // Intersection Observer for scroll animations
        const observerOptions = {
            threshold: 0.1,
            rootMargin: '0px 0px -50px 0px'
        };

        const observer = new IntersectionObserver((entries) => {
            entries.forEach(entry => {
                if (entry.isIntersecting) {
                    entry.target.style.opacity = '1';
                    entry.target.style.transform = 'translateY(0)';
                }
            });
        }, observerOptions);

        // Observe all cards
        document.querySelectorAll('.stat-card, .feature-card, .usecase-card, .demo-card, .example-card').forEach(card => {
            card.style.opacity = '0';
            card.style.transform = 'translateY(20px)';
            card.style.transition = 'opacity 0.6s ease-out, transform 0.6s ease-out';
            observer.observe(card);
        });
    }

    // Utility methods
    formatNumber(num) {
        return new Intl.NumberFormat().format(num);
    }

    formatBytes(bytes) {
        const sizes = ['Bytes', 'KB', 'MB', 'GB'];
        if (bytes === 0) return '0 Bytes';
        const i = Math.floor(Math.log(bytes) / Math.log(1024));
        return Math.round(bytes / Math.pow(1024, i) * 100) / 100 + ' ' + sizes[i];
    }

    formatDuration(ms) {
        if (ms < 1000) return ms + 'ms';
        if (ms < 60000) return (ms / 1000).toFixed(1) + 's';
        return (ms / 60000).toFixed(1) + 'm';
    }
}

// Initialize the app when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    new GoEducationalApp();
});

// Add some interactive features
document.addEventListener('DOMContentLoaded', () => {
    // Add hover effects to cards
    const cards = document.querySelectorAll('.stat-card, .feature-card, .usecase-card, .demo-card, .example-card');
    
    cards.forEach(card => {
        card.addEventListener('mouseenter', () => {
            card.style.transform = 'translateY(-5px) scale(1.02)';
        });
        
        card.addEventListener('mouseleave', () => {
            card.style.transform = 'translateY(0) scale(1)';
        });
    });

    // Add click effects to buttons
    const buttons = document.querySelectorAll('.nav-btn, .demo-btn');
    
    buttons.forEach(btn => {
        btn.addEventListener('click', () => {
            btn.style.transform = 'scale(0.95)';
            setTimeout(() => {
                btn.style.transform = 'scale(1)';
            }, 150);
        });
    });

    // Add keyboard navigation
    document.addEventListener('keydown', (e) => {
        if (e.key === 'ArrowLeft' || e.key === 'ArrowRight') {
            const navButtons = document.querySelectorAll('.nav-btn');
            const currentIndex = Array.from(navButtons).findIndex(btn => btn.classList.contains('active'));
            
            let newIndex;
            if (e.key === 'ArrowLeft') {
                newIndex = currentIndex > 0 ? currentIndex - 1 : navButtons.length - 1;
            } else {
                newIndex = currentIndex < navButtons.length - 1 ? currentIndex + 1 : 0;
            }
            
            navButtons[newIndex].click();
        }
    });

    // Add smooth scrolling for anchor links
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function (e) {
            e.preventDefault();
            const target = document.querySelector(this.getAttribute('href'));
            if (target) {
                target.scrollIntoView({
                    behavior: 'smooth',
                    block: 'start'
                });
            }
        });
    });
});
