document.addEventListener('DOMContentLoaded', () => {
    initNavigation();
    initTheme();
    loadDashboard();
    initCBT();
    initDiary();
    initSymptoms();
    initPsychology();
    initAffirmations();
});

function initNavigation() {
    document.querySelectorAll('.nav-item').forEach(item => {
        item.addEventListener('click', (e) => {
            e.preventDefault();
            const page = item.dataset.page;
            document.querySelectorAll('.page').forEach(p => p.classList.remove('active'));
            document.querySelectorAll('.nav-item').forEach(n => n.classList.remove('active'));
            document.getElementById(`${page}-page`).classList.add('active');
            item.classList.add('active');
            if (page === 'dashboard') loadDashboard();
            if (page === 'diary') loadDiary();
            if (page === 'symptoms') loadSymptomsList();
            if (page === 'psychology') loadPsychology();
        });
    });
}

function initTheme() {
    const toggle = document.getElementById('themeToggle');
    toggle?.addEventListener('click', () => {
        document.body.classList.toggle('dark');
        const icon = toggle.querySelector('i');
        icon.classList.toggle('fa-moon');
        icon.classList.toggle('fa-sun');
    });
}

async function loadDashboard() {
    try {
        const stats = await api.request('/stats');
        document.getElementById('stat-cbt').innerText = stats.cbt_count || 0;
        document.getElementById('stat-diary').innerText = stats.diary_count || 0;
    } catch(e) { console.error(e); }
}

function initCBT() {
    const form = document.getElementById('cbt-form');
    form?.addEventListener('submit', async (e) => {
        e.preventDefault();
        const thought = document.getElementById('cbt-thought').value;
        if (!thought) return;
        const result = await api.analyzeThought(thought);
        const container = document.getElementById('cbt-analysis-result');
        container.innerHTML = `
            <div class="analysis-result">
                <h4>Когнитивные искажения</h4>
                <div>${result.distortions?.map(d => `<span class="badge">${d}</span>`).join('') || 'Не обнаружено'}</div>
                <h4>Рациональный ответ</h4>
                <p>${result.rational_response || 'Попробуйте посмотреть иначе'}</p>
            </div>
        `;
    });
}

async function initDiary() {
    window.loadDiary = async () => {
        const sections = await api.getDiarySections();
        const container = document.getElementById('diary-sections');
        if (!container) return;
        container.innerHTML = sections.map(s => `
            <div class="card">
                <h3>${s.title}</h3>
                <p>${s.description}</p>
                <textarea id="diary-${s.id}" rows="4" placeholder="Напишите ответ..."></textarea>
                <button class="btn btn-sm" onclick="saveDiary('${s.id}')">Сохранить</button>
            </div>
        `).join('');
    };
    window.saveDiary = async (section) => {
        const answer = document.getElementById(`diary-${section}`).value;
        if (!answer) return;
        await api.saveDiaryEntry('current-user', section, answer, []);
        alert('Сохранено!');
    };
    loadDiary();
}

async function initSymptoms() {
    window.loadSymptomsList = async () => {
        const symptoms = await api.getSymptoms();
        const container = document.getElementById('symptoms-list');
        if (!container) return;
        container.innerHTML = symptoms.map(s => `
            <label class="symptom-item">
                <input type="checkbox" value="${s.id}"> ${s.name}
                <small>${s.category}</small>
            </label>
        `).join('');
    };
    document.getElementById('analyze-symptoms')?.addEventListener('click', async () => {
        const selected = [...document.querySelectorAll('#symptoms-list input:checked')].map(cb => cb.value);
        const result = await api.analyzeSymptoms(selected);
        document.getElementById('symptoms-result').innerHTML = `<div class="card"><h3>Результат</h3><p>${result.interpretation || 'Анализ завершен'}</p></div>`;
    });
    loadSymptomsList();
}

async function initPsychology() {
    window.loadPsychology = async () => {
        const traumas = await api.getTraumas();
        const container = document.getElementById('traumas-list');
        if (!container) return;
        container.innerHTML = Object.values(traumas).map(t => `
            <div class="card">
                <h3>${t.name_ru}</h3>
                <p>${t.description}</p>
                <h4>Симптомы:</h4>
                <ul>${t.symptoms?.map(s => `<li>${s}</li>`).join('') || ''}</ul>
                <h4>Пути исцеления:</h4>
                <ul>${t.healing_path?.map(h => `<li>${h}</li>`).join('') || ''}</ul>
            </div>
        `).join('');
    };
    loadPsychology();
}

async function initAffirmations() {
    window.refreshAffirmation = async () => {
        const affirmations = await api.getAffirmations('');
        if (affirmations.length) {
            document.getElementById('affirmation-text').innerText = affirmations[0].text;
        }
    };
    window.searchAffirmations = async () => {
        const keyword = document.getElementById('affirmation-search')?.value;
        if (!keyword) return;
        const results = await api.getAffirmations(keyword);
        const container = document.getElementById('affirmation-results');
        container.innerHTML = results.map(a => `<div class="affirmation-item"><p>${a.text}</p><small>${a.author}</small></div>`).join('');
    };
    refreshAffirmation();
}
