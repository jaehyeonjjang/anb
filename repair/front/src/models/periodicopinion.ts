import request from '~/global/request'

export default class Periodicopinion {
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/periodicopinion',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/periodicopinion',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodicopinion',
            data: item
        })

        return res
    }

    static async removebatch(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodicopinion/batch',
            data: item
        })

        return res
    }    

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/periodicopinion',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodicopinion/${id}`
        })

        return res
    }

    static async getByPeriodic(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodicopinion/get/periodic/${id}`
        })

        return res
    }
}
