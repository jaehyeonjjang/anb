import request from '~/global/request'

export default class Standard {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/standard',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/standard',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/standard',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/standard',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/standard/${id}`
        })

        return res
    }

    static async countByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/standard/count/apt/${apt}`
        })

        return res.count
    }

    static async findByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/standard/find/apt/${apt}`
        })

        return res
    }

    static async all(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/standard/all',
            data: item
        })

        return res
    }

    static async updateall(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/standard/all',
            data: item
        })

        return res
    }
}
