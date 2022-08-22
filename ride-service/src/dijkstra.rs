
use crate::models::*;

type Vertex = i32;

#[derive(Debug)]
pub struct Connection {
    pub peers   : (Vertex, Vertex),
    pub weight  : u64,
}

#[derive(Debug)]
pub struct Graph {
    pub connections : Vec<Connection>,
    vertices    : Vec<Vertex>,
}

#[derive(Debug)]
struct Road {
    vertex     : Vertex,
    distance   : u64,
    via_vertex : Vertex,
}

#[derive(Debug)]
pub struct DijkstraTable {
    start_vertex : Vertex,
    roads        : Vec<Road>,
    unvisited    : Vec<Vertex>,
}

impl DijkstraTable {
    fn get_distance(&self, vertex: Vertex) -> u64 {
        let mut ret = 0;

        for r in &self.roads {
            if r.vertex == vertex {
                ret = r.distance;
            }
        }

        ret
    }

    fn get_next_unvisited(&self) -> Option<Vertex> {
        let mut min = u64::MAX;
        let mut next = None;

        for vertex in &self.unvisited {
            for r in &self.roads
            {
              if r.vertex == *vertex 
              {
                if r.distance < min {
                  min = r.distance;
                  next = Some(*vertex);
                }
              }
            }
        }
        next
    }

    fn remove(&mut self, v : Vertex) {
        let mut index = 0;
        while index < self.unvisited.len() {
            let toremove = self.unvisited[index];
            if v == toremove {
                self.unvisited.remove(index);
                break
            }
            index += 1;
        }
    }


    pub fn new( graph: &Graph, start: Vertex ) -> DijkstraTable {
        let mut table = DijkstraTable {
            start_vertex : start,
            roads        : Vec::new(),
            unvisited    : graph.vertices.clone(),
        };

        for v in &graph.vertices {
            let mut road = Road::new(*v);

            if *v == start {
                road.distance = 0;
            }

            table.roads.push(road);
        }

        loop {
            match table.get_next_unvisited() {
                None => break,
                Some(v) => {
                    //println!("{}##################",v);

                    for n in graph.get_neighbours(v) {
                      let d = graph.get_weight((v, n));
                      let k = d + table.get_distance(v);
                      for road in &mut table.roads
                      {         
                        if road.vertex == n
                        {
                          if k < road.distance {
                            road.via_vertex = v;
                            road.distance = k;
                          }
                          break;
                        }
                      }
                    }
                    table.remove(v);
                    // println!(" {:#?} ", table);
                }
            }
            
        }
        table
    }
}

impl Road {
    fn new(from: Vertex) -> Road {
        Road {
            vertex      : from,
            distance    : u64::MAX,
            via_vertex  : 0,
        }
    }
}

impl Graph {
    fn get_weight(&self, peers: (Vertex, Vertex)) -> u64 {
        let mut ret : u64 = 0;

        for c in &self.connections {
            let (a, b) = peers;

            if c.peers == peers || c.peers == (b, a) {
                ret = c.weight;
                break;
            }
        }
        ret
    }

    fn get_neighbours(&self, vertex: Vertex) -> Vec<Vertex> {
        let mut neighbours : Vec<Vertex> = Vec::new();

        for c in &self.connections {
            if c.peers.0 == vertex {
                println!("for {} - adding to array: {}", vertex, c.peers.1);
                neighbours.push(c.peers.1);
            } // else if c.peers.1 == vertex {
            //    println!("for {} - adding to array: {}", vertex, c.peers.0);
            //    neighbours.push(c.peers.0);
            //}
        }

        neighbours
    }

    fn vertices_from_connections(conns : &Vec<Connection>) -> Vec<Vertex> {
        let mut verts : Vec<Vertex> = Vec::new();

        for c in conns.iter() {
            if ! verts.contains(&c.peers.0) {
                verts.push(c.peers.0);
            }
            if ! verts.contains(&c.peers.1) {
                verts.push(c.peers.1);
            }
        }
        verts
    }

    pub fn new(conns: Vec<Connection>) -> Graph {
        Graph {
            vertices    : Graph::vertices_from_connections(&conns),
            connections : conns,
        }
    }
}

pub fn dijkstra_fun() {
    let graph = Graph::new(
        vec![
            Connection {
                peers: (1, 3),
                weight: 31525200070,
            },
            Connection {
                peers: (1, 6),
                weight: 31525200070,
            },
            Connection {
                peers: (2, 3),
                weight: 31520700350,
            },
            Connection {
                peers: (2, 6),
                weight: 31520700350,
            },
            Connection {
                peers: (3, 4),
                weight: 31532400090,
            },
            Connection {
                peers: (4, 5),
                weight: 31534200140,
            },
            Connection {
                peers: (6, 7),
                weight: 31532400100,
            },
        ]
    );
    for connection in &graph.connections {
        println!("{} to {} - weight = {}", connection.peers.0, connection.peers.1, connection.weight);
    } 

    let dt = DijkstraTable::new( &graph, 3 );
    println!(" Dijkstra of '3': {:#?}", dt );
}


fn reconstruct_path(table: &DijkstraTable, begin_vertex: Vertex, end_vertex: Vertex) -> Vec<Vertex> {
    let mut path: Vec<Vertex> = Vec::new();
    let mut via: Vertex = end_vertex;

    while via != begin_vertex {
        path.push(via);
        via = find_via_vertex(table, via);
        if via == -1 {
            return Vec::new();
        }
    }   

    path
}

fn find_via_vertex(table: &DijkstraTable, vertex: Vertex) -> Vertex {
    let mut via: Vertex = -1;

    for road in &table.roads {
        if road.vertex == vertex {
            via = road.via_vertex;
        } 
    }

    via
}

fn find_drive_index_in_drives(drive_id: i32, drives: &Vec<Drive>) -> usize {
    let mut idx: usize  = 100000;

    for i in 0..drives.len() {
        if drives[i].id == drive_id {
            idx = i;
        }
    }

    idx
}

fn calc_path_size(path: &Vec<Vertex>, drives: &Vec<Drive>) -> u64 {
    let mut sum  = 0;
    
    for vertex in path {
        let drive_idx = find_drive_index_in_drives(*vertex, drives);

        sum += drives[drive_idx].distance as u64;
    }

    sum
}

pub fn dijkstra_fun_graph(graphh: &Graph, starting_ids: &Vec<Vertex>, ending_ids: &Vec<Vertex>, drives: &Vec<Drive>) {
    let mut conns: Vec<Connection> = Vec::new();
    
    let mut best_path:Vec<Vertex> = Vec::new();
    let mut best_path_size: u64 = std::u64::MAX;

    for connection in &graphh.connections {
        conns.push(Connection{
            peers: (connection.peers.0, connection.peers.1),
            weight: connection.weight,
        })
    } 

    let graph = Graph::new(conns);

    for starting_id in starting_ids {
        let dt = DijkstraTable::new( &graph, *starting_id);
        println!(" Dijkstra of {}: {:#?}", starting_id, dt);
        println!("\n");

        println!("{}", ending_ids.len());
        for ending_id in ending_ids {
            let mut path = reconstruct_path(&dt, *starting_id, *ending_id);
            
            if path.len() != 0 {
                path.push(*starting_id);
                let path_size = calc_path_size(&path, drives);
                println!("Path for starting id - {}, ending id - {}, size - {}", starting_id, ending_id, path_size);
                
                for vertex in &path {
                    println!("{}", vertex);
                }

                if best_path_size > path_size {
                    best_path_size = path_size;
                    best_path = path.clone();
                }
            
            } else {
                println!("There is no path from {} to {}", starting_id, ending_id);
            }
            
            
            println!();
        }

        println!("---------------------------------------------\n\n");
    }
    
    println!("--------Shortest path-----------");
    for vertex in best_path.iter().rev() {
        println!("{}", vertex)
    }
    println!("--------             -----------");

}