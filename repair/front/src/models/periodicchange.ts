import request from '~/global/request'

export default class Periodicchange {
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/periodicchange',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/periodicchange',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodicchange',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/periodicchange',
            params: params
        })

        if (res.items == null) {
            res.items = []
        }

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodicchange/${id}`
        })

        return res
    }
}
