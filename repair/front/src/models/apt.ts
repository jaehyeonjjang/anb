import request from '~/global/request'

export default class Apt {
    static contracttypes = ['계약구분', '장기수선계획', '정밀', '정기', '하자보수', '하자조사', '정밀안전진단', '감리', '기술자문']
    static contracttypeTypes = ['', 'info', 'success', 'warning', 'danger']

    static invoices = ['계산서 발행여부', '발행', '미발행']
    static invoiceTypes = ['', 'warning', 'info']    

    static getContracttypeType(value: number) {
        return this.contracttypeTypes[value]
    }

    static getContracttype(value: number) {
        return this.contracttypes[value]
    }

    static getInvoiceType(value: number) {
        return this.invoiceTypes[value]
    }

    static getInvoice(value: number) {
        return this.invoices[value]
    }
    
    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/apt',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/apt',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/apt',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/apt',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/apt/${id}`
        })

        return res
    }

    static async search(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/apt/search',
            params: params
        })

        return res
    }
}
