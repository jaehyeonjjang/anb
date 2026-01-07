import request from '~/global/request'

export default class Standardlist {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/standardlist',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/standardlist',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/standardlist',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/standardlist',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/standardlist/${id}`
        })

        return res
    }

    static async countByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/standardlist/count/apt/${apt}`
        })

        return res.count
    }

    static async findByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/standardlist/find/apt/${apt}`
        })

        return res
    }        
}
