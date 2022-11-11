package day6

import "strings"

type node struct {
	// 待匹配路由，/p/:lang
	urlPath string
	// 路由中的一部分，:lang
	part string
	// 子节点，[doc, tutorial, intro]
	children []*node
	// 是否模糊匹配，part 含有 : 或 * 时为true
	isDynamic bool
}

// 找到第一个匹配成功的节点，用于插入
func (currentNode *node) matchChild(part string) *node {
	for _, child := range currentNode.children {
		// 首先查找子节点中是否有该part路径, 或者子节点中是否有动态路由
		if child.part == part || child.isDynamic {
			// 匹配成功
			return child
		}
	}
	return nil
}

// 查找所有匹配的节点，用于查找
func (currentNode *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range currentNode.children {
		if child.part == part || child.isDynamic {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 开发时，注册路由规则，映射handler；
// 访问时，匹配路由规则，查找对应的handler；

// 插入: 注册路由，映射handler
// 存在的问题，出现路由冲突时，会覆盖之前的路由（待解决）
func (currentNode *node) insert(urlPath string, parts []string, height int) {
	// 退出递归的条件
	if len(parts) == height {
		currentNode.urlPath = urlPath
		return
	}

	part := parts[height]
	child := currentNode.matchChild(part)

	// 若没有匹配的子路径节点，
	// 则创建一个节点并添加到当前节点的子节点中
	if child == nil {
		child = &node{
			part:      part,
			isDynamic: part[0] == ':' || part[0] == '*'}
		// node的children字段是一个node指针切片
		currentNode.children = append(currentNode.children, child)
	}
	// 递归插入
	child.insert(urlPath, parts, height+1)
}

// 查找：匹配路由，查找对应的handler
func (currentNode *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(currentNode.part, "*") {
		if currentNode.urlPath == "" {
			return nil
		}
		return currentNode
	}

	part := parts[height]
	children := currentNode.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)

		if result != nil {
			return result
		}
	}

	return nil
}
