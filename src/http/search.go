// Copyright 2017 Kevin Bayes
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package http

type ListMeta struct {

	Size int
	Page int
	Pages int
}

type ResponseList struct {
	Meta ListMeta
	Entities interface{}
	Links []Link
}

func MakeSearchResult(size int, pages int, page int, entities interface{}, links []Link) *ResponseList {

	return &ResponseList{
		Meta: ListMeta{
			Size: size,
			Page: page,
			Pages: pages,
		},
		Entities: entities,
		Links: links,
	}
}