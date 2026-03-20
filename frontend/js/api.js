const API_BASE = 'http://localhost:8082/api/v1';

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
        return this.request('/cbt/analyze', { 
            method: 'POST', 
            body: JSON.stringify({ thought }) 
        });
    }
    
    async getSymptoms() {
        return this.request('/symptoms');
    }
    
    async getAffirmations(keyword) {
        return this.request(`/affirmations?keyword=${encodeURIComponent(keyword)}`);
    }
    
    async getTraumas() {
        return this.request('/psychology/traumas');
    }
    
    async getDiarySections() {
        return this.request('/diary/sections');
    }
    
    async getDiaryQuestions() {
        return this.request('/diary/questions');
    }
    
    async saveDiaryEntry(userId, section, answer, tags) {
        return this.request('/diary/entries', {
            method: 'POST',
            body: JSON.stringify({ userId, section, answer, tags })
        });
    }
    
    async getStats() {
        return this.request('/stats');
    }
}

const api = new API();
