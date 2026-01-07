import request from '~/global/request'

export default class Periodic {
    static statuss = ['상태', '준비', '착수', '완료', '중단']
    static statusTypes = ['', 'info', 'success', 'warning', 'danger']

    static getStatusType(value: number) {
        return this.statusTypes[value]
    }

    static getStatus(value: number) {
        return this.statuss[value]
    }
    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/periodic',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/periodic',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodic',
            data: item
        })

        return res
    }

    static async removebatch(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodic/batch',
            data: item
        })

        return res
    }    

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/periodic',
            params: params
        })

        return res
    }

    static async search(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/periodic/search',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodic/${id}`
        })

        return res
    }

    static async duplication(id: number) {
        let item = {
            id
        }
        
        const res = await request({
            method: 'POST',
            url: '/api/periodic/duplication',
            data: item
        })

        return res
    }    
}
