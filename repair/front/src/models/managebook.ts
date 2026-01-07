import request from '~/global/request'

export default class Managebook {
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/managebook',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/managebook',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/managebook',
            data: item
        })

        return res
    }

    static async removebatch(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/managebook/batch',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/managebook',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/managebook/${id}`
        })

        return res
    }

    static async process(id: number, name: string, order: number, filename: string) {
        const item = {
            id: id,
            name: name,
            order: order,
            filename: filename
        }
        
        const res = await request({
            method: 'POST',
            url: `/api/managebook/process`,
            data: item
        })

        return res
    }

    static async multiprocess(id: number, filename: string, originalfilename: string) {
        const item = {
            id: id,            
            filename: filename,
            originalfilename: originalfilename
        }

        const res = await request({
            method: 'POST',
            url: `/api/managebook/multiprocess`,
            data: item
        })

        return res
    }    
}
