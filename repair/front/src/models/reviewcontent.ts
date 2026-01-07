import request from '~/global/request'

export default class Reviewcontent {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/reviewcontent',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/reviewcontent',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/reviewcontent',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/reviewcontent',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/reviewcontent/${id}`
        })

        return res
    }

    static async countByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/reviewcontent/count/apt/${apt}`
        })

        return res.count
    }

    static async findByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/reviewcontent/find/apt/${apt}`
        })

        return res
    }        
}
