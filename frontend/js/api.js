const API_BASE = '/api/v1';

class API {
    async request(endpoint, options = {}) {
        const response = await fetch(`${API_BASE}${endpoint}`, {
            headers: { 'Content-Type': 'application/json' },
            ...options
        });
        if (!response.ok) throw new Error(await response.text());
        return response.json();
    }
    async analyzeThought(thought) {
        return this.request('/cbt/analyze', { method: 'POST', body: JSON.stringify({ thought }) });
    }
    async getDiarySections() {
        return this.request('/diary/sections');
    }
    async getDiaryQuestions() {
        return this.request('/diary/questions');
    }
    async saveDiaryEntry(userId, section, answer, tags) {
        return this.request('/diary/entries', { method: 'POST', body: JSON.stringify({ userId, section, answer, tags }) });
    }
    async getSymptoms() {
        return this.request('/symptoms');
    }
    async analyzeSymptoms(symptomIds) {
        return this.request('/symptoms/analyze', { method: 'POST', body: JSON.stringify({ symptom_ids: symptomIds }) });
    }
    async getTraumas() {
        return this.request('/psychology/traumas');
    }
    async getAffirmations(keyword) {
        return this.request(`/affirmations?keyword=${encodeURIComponent(keyword)}`);
    }
}

const api = new API();
