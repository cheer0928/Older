import axios from 'axios'
import type { NursingHome } from '@/types/nursing-home'

const api = axios.create({
    baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api'
})

export const nursingHomeApi = {
    getAll: () => api.get<NursingHome[]>('/nursing-homes'),
    getById: (id: number) => api.get<NursingHome>(`/nursing-homes/${id}`),
    create: (data: Partial<NursingHome>) => api.post('/nursing-homes', data),
    update: (id: number, data: Partial<NursingHome>) => api.put(`/nursing-homes/${id}`, data),
    delete: (id: number) => api.delete(`/nursing-homes/${id}`)
} 