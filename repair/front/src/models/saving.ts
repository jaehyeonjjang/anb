import request from '~/global/request'

export default class Saving {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/saving',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/saving',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/saving',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/saving',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/saving/${id}`
        })

        return res
    }

    static async countByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/saving/count/apt/${apt}`
        })

        return res.count
    }

    static async findByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/saving/find/apt/${apt}`
        })

        return res
    }        
}
