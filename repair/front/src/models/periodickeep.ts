import request from '~/global/request'

export default class Periodickeep {
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/periodickeep',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/periodickeep',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodickeep',
            data: item
        })

        return res
    }

    static async removebatch(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodickeep/batch',
            data: item
        })

        return res
    }    

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/periodickeep',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodickeep/${id}`
        })

        return res
    }

    static async getByPeriodic(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodickeep/get/periodic/${id}`
        })

        return res
    }
}
