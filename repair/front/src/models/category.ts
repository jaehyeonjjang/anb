import request from '~/global/request'

export default class Category {    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/category',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/category',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/category',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/category',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/category/${id}`
        })

        return res
    }

    static async countByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/category/count/apt/${apt}`
        })

        return res.count
    }

    static async findByApt(apt: number) {
        const res = await request({
            method: 'GET',
            url: `/api/category/find/apt/${apt}`
        })

        return res
    }

    static async countByAptLevel(apt: number, level: number) {
        const res = await request({
            method: 'GET',
            url: `/api/category/count/aptlevel/${apt}?level=${level}`
        })

        return res.count
    }

    static async findByAptLevel(apt: number, level: number) {
        const res = await request({
            method: 'GET',
            url: `/api/category/find/aptlevel/${apt}?level=${level}`
        })

        return res
    }    

    static async init(apt: number) {
        let item = {
            apt
        }
        
        const res = await request({
            method: 'POST',
            url: '/api/category/initdata',
            data: item
        })

        return res
    }

    static async duplication(apt: number) {
        let item = {
            apt
        }
        
        const res = await request({
            method: 'POST',
            url: '/api/category/duplicationdata',
            data: item
        })

        return res
    }    
    
}
